/**
 * MIT License
 *
 * Copyright (c) 2021 clevabit GmbH
 * Copyright (c) 2021 trilogik GmbH
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"fmt"
	"github.com/clbanning/mxj/v2"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("test-atmega.cproj")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	mxj.XmlCharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "ASCII" {
			return input, nil
		}
		return nil, fmt.Errorf("illegal charset: %s", charset)
	}
	elements, err := mxj.NewMapXml(content)
	if err != nil {
		panic(err)
	}

	nodes := make([]*Element, 0)
	n, err := parseElements(elements)
	if err != nil {
		panic(err)
	}
	nodes = append(nodes, n...)

	f, err = os.OpenFile("types/generated2.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if _, err := f.WriteString("package dfp\n\n"); err != nil {
		panic(err)
	}

	types := make([]*Element, 0)
	for _, element := range nodes {
		types = collectTypes(element, types)
	}

	for _, element := range types {
		types = writeType(element, types, f)
	}

	f.Close()
}

func collectTypes(element *Element, types []*Element) []*Element {
	name := toCamelCase(element.name)
	if _, found := findElement(types, name); !found {
		types = append(types, element)
	}
	for _, child := range element.children {
		types = collectTypes(child, types)
	}
	return types
}

func writeType(element *Element, types []*Element, f io.StringWriter) []*Element {
	name := toTypeName(element.name)
	if _, err := f.WriteString(fmt.Sprintf("type %s struct {\n", name)); err != nil {
		panic(err)
	}

	for _, attribute := range element.attributes {
		attName := toTypeName(attribute.name)
		datatype := "string"
		if attribute.datatype != nil {
			switch *attribute.datatype {
			case integer:
				datatype = "int64"
			case number:
				datatype = "float64"
			case boolean:
				datatype = "bool"
			}
		}
		if _, err := f.WriteString(fmt.Sprintf("\t%s %s `xml:\"%s,attr\"`\n", attName, datatype, attribute.name)); err != nil {
			panic(err)
		}
	}

	for _, child := range element.children {
		childName := toTypeName(child.name)
		propertyName := childName
		array := ""
		if child.array {
			array = "[]"
			if !strings.HasSuffix(propertyName, "s") {
				propertyName += "s"
			}
		}
		if _, err := f.WriteString(fmt.Sprintf("\t%s %s%s `xml:\"%s\"`\n", propertyName, array, childName, child.name)); err != nil {
			panic(err)
		}
		if _, found := findElement(types, child.name); !found {
			types = append(types, child)
		}
	}

	if element.inline {
		datatype := "string"
		if element.datatype != nil {
			switch *element.datatype {
			case integer:
				datatype = "int64"
			case number:
				datatype = "float64"
			case boolean:
				datatype = "bool"
			}
		}
		if _, err := f.WriteString(fmt.Sprintf("\tContent %s `xml:\",chardata\"`\n", datatype)); err != nil {
			panic(err)
		}
	}

	if _, err := f.WriteString("}\n\n"); err != nil {
		panic(err)
	}

	return types
}

func toTypeName(name string) string {
	name = strings.Replace(name, ".", "-", -1)
	name = strings.Replace(name, "_", "-", -1)
	typeName := ""
	tokens := strings.Split(name, "-")
	for _, token := range tokens {
		typeName = typeName + toCamelCase(token)
	}
	return typeName
}

func toCamelCase(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func parseElements(elements map[string]interface{}) ([]*Element, error) {
	nodes := make([]*Element, 0)
	for key, value := range elements {
		if _, found := findElement(nodes, key); !found {
			node, err := parseNode(key, value, nil)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func parseNode(key string, value interface{}, parent *Element) (*Element, error) {
	if parent == nil {
		parent = &Element{
			name: key,
		}
	}

	switch v := value.(type) {
	case map[string]interface{}:
		if err := parseElement(parent, v); err != nil {
			return nil, err
		}

	case []interface{}:
		parent.array = true
		for _, el := range v {
			if _, err := parseNode(key, el, parent); err != nil {
				return nil, err
			}
		}

	case string:
		// inline value
		parent.inline = true
		parent.datatype = guessDatatype(v, parent.datatype)

	default:
		fmt.Printf("BROKEN: %+v\n", v)
	}
	return parent, nil
}

func parseElement(node *Element, element map[string]interface{}) error {
	for key, value := range element {
		if strings.HasPrefix(key, "-") {
			key = key[1:]
			if attribute, found := findAttribute(node.attributes, key); !found {
				node.attributes = append(node.attributes, &Attribute{
					name:     key,
					datatype: guessDatatype(value.(string), nil),
				})
			} else {
				attribute.datatype = guessDatatype(value.(string), attribute.datatype)
			}
			continue
		} else if key == "#text" {
			// inline value
			node.inline = true
			node.datatype = guessDatatype(value.(string), node.datatype)
			continue
		}

		existing, found := findElement(node.children, key)
		child, err := parseNode(key, value, existing)
		if err != nil {
			return err
		}
		if !found {
			node.children = append(node.children, child)
		}
	}
	return nil
}

func findElement(nodes []*Element, name string) (*Element, bool) {
	for _, node := range nodes {
		if node.name == name {
			return node, true
		}
	}
	return nil, false
}

func findAttribute(attributes []*Attribute, name string) (*Attribute, bool) {
	for _, attribute := range attributes {
		if attribute.name == name {
			return attribute, true
		}
	}
	return nil, false
}

func guessDatatype(content string, currentGuess *Datatype) *Datatype {
	newGuess := guesser(content)
	if currentGuess != nil {
		if *currentGuess == text {
			return currentGuess
		}
		if *currentGuess == newGuess {
			return currentGuess
		}
		if newGuess == text {
			return &newGuess
		}
		if newGuess == number && *currentGuess == integer {
			return &newGuess
		}
		if newGuess != *currentGuess {
			val := text
			return &val
		}
	}
	return &newGuess
}

func guesser(content string) Datatype {
	if _, err := strconv.ParseInt(content, 10, 64); err == nil {
		return integer
	}
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		return number
	}
	if _, err := strconv.ParseBool(strings.ToLower(content)); err == nil {
		return boolean
	}
	return text
}

type Attribute struct {
	name     string
	datatype *Datatype
}

type Element struct {
	name       string
	inline     bool
	array      bool
	datatype   *Datatype
	children   []*Element
	attributes []*Attribute
}

type Datatype int

const (
	text Datatype = iota
	integer
	number
	boolean
)
