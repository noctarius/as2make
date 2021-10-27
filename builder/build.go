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

package builder

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/clevabit/as2make/cmsis"
	"github.com/clevabit/as2make/dfp"
	"github.com/clevabit/as2make/toolchain"
	"github.com/clevabit/as2make/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	project         = flag.String("project", "-", "--project=<path_to_project_file>")
	toolchainPath   = flag.String("toolchain", "-", "--toolchain=<path_to_bin_dir>")
	toolchainPrefix = flag.String("toolchain-prefix", "-", "--toolchain-prefix=<CROSS_COMPILE>")
	cmsisSdkPath    = flag.String("cmsis-sdk", "", "--cmsis-sdk=<path_to_cmsis_dir>")
	dfpSdkPath      = flag.String("dfp-sdk", "-", "--dfp-sdk=<path_to_dfp_dir>")
	outputName      = flag.String("output-name", "build", "--output-name=<outputname>")
)

func init() {
	flag.Parse()
}

type Build struct {
	projectDir               string
	project                  types.Project
	toolchain                toolchain.Toolchain
	sources                  []File
	ldSources                []File
	subdirs                  []string
	settings                 types.PropertyGroup
	release                  types.PropertyGroup
	includePaths             []string
	defSymbols               []string
	linkerLibrarySearchPaths []string
	cmsis                    cmsis.SDK
	dfp                      dfp.SDK
	optimizationLevel        string
	deviceDefine             string
	coreSpecification        string
	outputName               string
}

func NewBuild() (Build, error) {
	toolchain, err := toolchain.NewToolchain(*toolchainPath, *toolchainPrefix)
	if err != nil {
		panic(err)
	}

	if *project == "-" {
		return Build{}, fmt.Errorf("no valid project file given")
	}

	projectDir := filepath.Dir(*project)
	project, err := loadProject(*project)
	if err != nil {
		return Build{}, err
	}

	sources, err := findSources(project, projectDir)
	if err != nil {
		return Build{}, err
	}

	ldSources, err := findLinkerSources(project, projectDir)
	if err != nil {
		return Build{}, err
	}

	subdirs, err := findSubdirs(project, projectDir)
	if err != nil {
		return Build{}, err
	}

	release, found := findRelease(project)
	if !found {
		return Build{}, fmt.Errorf("no release configuration found")
	}

	if release.ToolchainSettings.AvrGcc == nil && release.ToolchainSettings.ArmGcc == nil {
		return Build{}, fmt.Errorf("found unsupported cproj file (at the moment only avr and arm are supported)")
	}

	settings, found := findProjectSettings(project)
	if !found {
		return Build{}, fmt.Errorf("no project settings configuration found")
	}

	cmsis, err := cmsis.New(*cmsisSdkPath)
	if err != nil {
		return Build{}, err
	}

	dfp, err := dfp.New(*dfpSdkPath)
	if err != nil {
		return Build{}, err
	}

	device, err := dfp.Device(settings.Avrdevice.Content)
	if err != nil {
		return Build{}, err
	}

	includePaths := includePaths(release.ToolchainSettings, cmsis)
	startupIncludePath, err := deviceSpecificIncludePath(settings, dfp)
	if err != nil {
		return Build{}, err
	}
	includePaths = append(includePaths, startupIncludePath)

	return Build{
		projectDir:               projectDir,
		project:                  project,
		toolchain:                toolchain,
		sources:                  sources,
		subdirs:                  subdirs,
		ldSources:                ldSources,
		settings:                 settings,
		release:                  release,
		includePaths:             includePaths,
		defSymbols:               defSymbols(release.ToolchainSettings),
		optimizationLevel:        optimizationLevel(release.ToolchainSettings),
		deviceDefine:             device.Compile.Define,
		coreSpecification:        device.Processor.Dcore,
		outputName:               *outputName,
		linkerLibrarySearchPaths: linkerLibrarySearchPaths(release.ToolchainSettings),
	}, nil
}

func (b Build) OutputName(ext string) string {
	if strings.HasPrefix(ext, ".") {
		ext = ext[1:]
	}
	return fmt.Sprintf("%s.%s", b.outputName, ext)
}

func (b Build) Sources() []File {
	return b.sources
}

func (b Build) Subdirs() []string {
	return b.subdirs
}

func (b Build) LinkerSources() []File {
	return b.ldSources
}

func (b Build) IncludePaths() []string {
	return b.includePaths
}

func (b Build) DefSymbols() []string {
	return b.defSymbols
}

func (b Build) Toolchain() toolchain.Toolchain {
	return b.toolchain
}

func (b Build) OptimizationLevel() string {
	return b.optimizationLevel
}

func (b Build) WithHex() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccCommonOutputfilesHex.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCommonOutputfilesHex.Content
}

func (b Build) WithLss() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccCommonOutputfilesLss.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCommonOutputfilesLss.Content
}

func (b Build) WithEep() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccCommonOutputfilesEep.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCommonOutputfilesEep.Content
}

func (b Build) WithBin() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return false
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCommonOutputfilesBin.Content
}

func (b Build) WithSrec() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccCommonOutputfilesSrec.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCommonOutputfilesSrec.Content
}

func (b Build) WithWarningAll() bool {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccCompilerWarningsAllWarnings.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccCompilerWarningsAllWarnings.Content
}

func (b Build) MiscellaneousLinkerFlags() string {
	if b.release.ToolchainSettings.AvrGcc != nil {
		return b.release.ToolchainSettings.AvrGcc.AvrgccLinkerMiscellaneousLinkerFlags.Content
	}
	return b.release.ToolchainSettings.ArmGcc.ArmgccLinkerMiscellaneousLinkerFlags.Content
}

func (b Build) DeviceDefine() string {
	return b.deviceDefine
}

func (b Build) CoreSpecification() string {
	core := strings.ToLower(b.coreSpecification)
	core = strings.Replace(core, "+", "plus", -1)
	return core
}

func (b Build) LinkerLibrarySearchPaths() []string {
	return b.linkerLibrarySearchPaths
}

func findSubdirs(project types.Project, projectDir string) ([]string, error) {
	folders := make([]string, 0)
	for _, itemGroup := range project.ItemGroups {
		if itemGroup.Folders != nil && len(itemGroup.Folders) > 0 {
			for _, folder := range itemGroup.Folders {
				candidate := filepath.Join(projectDir, strings.Replace(folder.Include, "\\", "/", -1))
				fi, _ := os.Stat(candidate)
				if fi != nil {
					folders = append(folders, strings.Replace(folder.Include, "\\", "/", -1))
				}
			}
		}
	}
	return folders, nil
}

func findSources(project types.Project, projectDir string) ([]File, error) {
	files := make([]File, 0)
	for _, itemGroup := range project.ItemGroups {
		if itemGroup.Compiles != nil && len(itemGroup.Compiles) > 0 {
			for _, compile := range itemGroup.Compiles {
				filename := compile.Include
				if strings.HasSuffix(filename, ".c") {
					filename = strings.Replace(filename, "\\", "/", -1)
					fi, _ := os.Stat(filename)
					if fi != nil {
						files = append(files, File{name: fmt.Sprintf("./%s", filename)})
					} else {
						// try absolute path
						filename = filepath.Join(projectDir, filename)
						fi, _ = os.Stat(filename)
						if fi != nil {
							files = append(files, File{name: filename})
						}
					}
				}
			}
		}
	}
	return files, nil
}

func findLinkerSources(project types.Project, projectDir string) ([]File, error) {
	files := make([]File, 0)
	for _, itemGroup := range project.ItemGroups {
		if itemGroup.Nones != nil && len(itemGroup.Nones) > 0 {
			for _, none := range itemGroup.Nones {
				filename := none.Include
				if strings.HasSuffix(filename, ".ld") {
					filename = strings.Replace(filename, "\\", "/", -1)
					fi, _ := os.Stat(filename)
					if fi != nil {
						files = append(files, File{name: filename})
					} else {
						// try absolute path
						filename = filepath.Join(projectDir, filename)
						fi, _ = os.Stat(filename)
						if fi != nil {
							files = append(files, File{name: filename})
						}
					}
				}
			}
		}
	}
	return files, nil
}

func loadProject(projectFile string) (types.Project, error) {
	fmt.Printf("Loading project file: %s ...\n", projectFile)

	f, err := os.Open(projectFile)
	if err != nil {
		return types.Project{}, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return types.Project{}, err
	}

	var project types.Project
	if err := xml.Unmarshal(content, &project); err != nil {
		return types.Project{}, err
	}
	return project, nil
}

func findRelease(project types.Project) (types.PropertyGroup, bool) {
	for _, propertyGroup := range project.PropertyGroups {
		if strings.Contains(propertyGroup.Condition, "Release") {
			return propertyGroup, true
		}
	}
	return types.PropertyGroup{}, false
}

func findProjectSettings(project types.Project) (types.PropertyGroup, bool) {
	for _, propertyGroup := range project.PropertyGroups {
		if propertyGroup.Condition == "" {
			return propertyGroup, true
		}
	}
	return types.PropertyGroup{}, false
}

func includePaths(toolchainSettings types.ToolchainSettings, cmsis cmsis.SDK) []string {
	var values types.ListValues
	if toolchainSettings.AvrGcc != nil {
		values = toolchainSettings.AvrGcc.AvrgccAssemblerGeneralIncludePaths.ListValues
	} else {
		values = toolchainSettings.ArmGcc.ArmgccAssemblerGeneralIncludePaths.ListValues
	}

	paths := make([]string, 0)
	path := strings.Replace(filepath.Join(cmsis.Path(), "/Core/Include"), "\\", "/", -1)
	paths = append(paths, fmt.Sprintf("-I\"%s\"", path))
	for _, value := range values.Values {
		if !strings.Contains(value.Content, "%24") {
			path = strings.Replace(value.Content, "\\", "/", -1)
			path = strings.Replace(path, "..", ".", -1)
			paths = append(paths, fmt.Sprintf("-I\"%s\"", path))
		}
	}
	return paths
}

func deviceSpecificIncludePath(settings types.PropertyGroup, dfp dfp.SDK) (string, error) {
	deviceName := settings.Avrdevice.Content
	component, err := dfp.Component("Device", "Startup", deviceName)
	if err != nil {
		return "", err
	}
	for _, file := range component.Files.Files {
		if file.Category == "include" {
			path := strings.Replace(filepath.Join(dfp.Path(), file.Name), "\\", "/", -1)
			return path, nil
		}
	}
	return "", fmt.Errorf("could not find device startup include")
}

func defSymbols(toolchainSettings types.ToolchainSettings) []string {
	var values types.ListValues
	if toolchainSettings.AvrGcc != nil {
		values = toolchainSettings.AvrGcc.AvrgccCompilerSymbolsDefSymbols.ListValues
	} else {
		values = toolchainSettings.ArmGcc.ArmgccCompilerSymbolsDefSymbols.ListValues
	}

	symbols := make([]string, 0)
	for _, value := range values.Values {
		symbols = append(symbols, fmt.Sprintf("-D%s", value.Content))
	}
	return symbols
}

func optimizationLevel(toolchainSettings types.ToolchainSettings) string {
	var level string
	if toolchainSettings.AvrGcc != nil {
		level = toolchainSettings.AvrGcc.AvrgccCompilerOptimizationLevel.Content
	} else {
		level = toolchainSettings.ArmGcc.ArmgccCompilerOptimizationLevel.Content
	}

	switch level {
	case "Optimize for size (-Os)":
		return "-Os"
	}
	return ""
}

func linkerLibrarySearchPaths(toolchainSettings types.ToolchainSettings) []string {
	var values types.ListValues
	if toolchainSettings.AvrGcc != nil {
		return []string{}
	} else {
		values = toolchainSettings.ArmGcc.ArmgccLinkerLibrariesLibrarySearchPaths.ListValues
	}

	paths := make([]string, 0)
	for _, value := range values.Values {
		path := strings.Replace(value.Content, "%24(ProjectDir)", ".", -1)
		if !strings.Contains(path, "%24") {
			path = strings.Replace(path, "\\", "/", -1)
			path = strings.Replace(path, "..", ".", -1)
			paths = append(paths, fmt.Sprintf("-L\"%s\"", path))
		}
	}
	return paths
}
