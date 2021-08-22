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
	"github.com/clevabit/as2make/builder"
	"github.com/clevabit/as2make/types"
	"io"
	"strings"
)

func main() {
	build, err := builder.NewBuild()
	if err != nil {
		panic(err)
	}

	writeMakefile(build)
	writeMakedepMk(build)
}

func searchPaths(toolchainSettings types.ToolchainSettings, cmsis, dfp string) []string {
	values := toolchainSettings.ArmGcc.ArmgccAssemblerGeneralIncludePaths.ListValues
	paths := make([]string, 0)
	for _, value := range values.Values {
		path := strings.Replace(value.Content, "%24(ProjectDir)", ".", -1)
		if !strings.Contains(value.Content, "%24") {
			path = strings.Replace(path, "\\", "/", -1)
			path = strings.Replace(path, "..", ".", -1)
			paths = append(paths, fmt.Sprintf("-I\"%s\"", path))
		}
	}
	return paths
}

func write(file io.StringWriter, format string, args ...interface{}) {
	if _, err := file.WriteString(fmt.Sprintf(format, args...)); err != nil {
		panic(err)
	}
}
