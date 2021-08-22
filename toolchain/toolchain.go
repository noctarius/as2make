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

package toolchain

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Toolchain struct {
	path   string
	prefix string
}

func NewToolchain(path, prefix string) (Toolchain, error) {
	prefix, err := findToolchainPrefix(prefix)
	if err != nil {
		return Toolchain{}, err
	}

	path = strings.Replace(path, "\\", "/", -1)
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	if !strings.HasSuffix(path, "/bin") {
		path = fmt.Sprintf("%s/bin", path)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return Toolchain{}, err
	}
	return Toolchain{
		path:   absPath,
		prefix: prefix,
	}, nil
}

func (t Toolchain) Executable(executable string) string {
	executable = fmt.Sprintf("%s%s", t.prefix, executable)
	return filepath.Join(t.path, executable)
}

func findToolchainPrefix(prefix string) (string, error) {
	cc, found := os.LookupEnv("CROSS_COMPILE")
	if !found && prefix == "-" {
		return "", fmt.Errorf("no valid toolchain prefix found")
	}
	if prefix != "-" {
		return prefix, nil
	}
	return cc, nil
}
