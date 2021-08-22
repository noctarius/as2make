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

package dfp

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type SDK struct {
	path        string
	packageDesc Package
}

func New(sdkPath string) (SDK, error) {
	packageDesc, err := loadPackageDescriptor(sdkPath)
	if err != nil {
		return SDK{}, nil
	}

	return SDK{
		path:        sdkPath,
		packageDesc: packageDesc,
	}, nil
}

func (s SDK) Path() string {
	return s.path
}

func (s SDK) Device(name string) (Device, error) {
	for _, device := range s.packageDesc.Devices.Family.Devices {
		if device.Dname == name {
			return device, nil
		}
	}
	return Device{}, fmt.Errorf("device %s not found in package descriptor", name)
}

func loadPackageDescriptor(sdkPath string) (Package, error) {
	files := make([]string, 0)
	if err := filepath.WalkDir(sdkPath, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".pdsc") {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return Package{}, err
	}

	if len(files) > 1 {
		return Package{}, fmt.Errorf("more than one package descriptor found")
	}
	if len(files) == 0 {
		return Package{}, fmt.Errorf("no package descriptor found in %s", sdkPath)
	}

	file := files[0]
	fmt.Printf("Loading DPF package descriptor: %s ...\n", file)

	f, err := os.Open(file)
	if err != nil {
		return Package{}, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return Package{}, err
	}

	var packageDesc Package
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	if err := decoder.Decode(&packageDesc); err != nil {
		return Package{}, err
	}
	return packageDesc, nil
}
