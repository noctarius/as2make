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

package types

type Project struct {
	Xmlns string `xml:"xmlns,attr"`
	ToolsVersion float64 `xml:"ToolsVersion,attr"`
	DefaultTargets string `xml:"DefaultTargets,attr"`
	PropertyGroups []PropertyGroup `xml:"PropertyGroup"`
	ItemGroups []ItemGroup `xml:"ItemGroup"`
	Import Import `xml:"Import"`
}

type PropertyGroup struct {
	Condition string `xml:"Condition,attr"`
	RootNamespace RootNamespace `xml:"RootNamespace"`
	BootSegment BootSegment `xml:"BootSegment"`
	AtStartFilePath AtStartFilePath `xml:"atStartFilePath"`
	AtzipPath AtzipPath `xml:"AtzipPath"`
	Avrtool Avrtool `xml:"avrtool"`
	ProjectVersion ProjectVersion `xml:"ProjectVersion"`
	ProgFlashFromRam ProgFlashFromRam `xml:"ProgFlashFromRam"`
	RamSnippetAddress RamSnippetAddress `xml:"RamSnippetAddress"`
	AsfFrameworkConfig AsfFrameworkConfig `xml:"AsfFrameworkConfig"`
	ComAtmelAvrdbgToolAtmelice ComAtmelAvrdbgToolAtmelice `xml:"com_atmel_avrdbg_tool_atmelice"`
	ComAtmelAvrdbgToolJtagice3plus ComAtmelAvrdbgToolJtagice3plus `xml:"com_atmel_avrdbg_tool_jtagice3plus"`
	ToolchainName ToolchainName `xml:"ToolchainName"`
	ProjectGuid ProjectGuid `xml:"ProjectGuid"`
	OutputFileExtension OutputFileExtension `xml:"OutputFileExtension"`
	Avrtoolserialnumber Avrtoolserialnumber `xml:"avrtoolserialnumber"`
	SchemaVersion SchemaVersion `xml:"SchemaVersion"`
	OutputFileName OutputFileName `xml:"OutputFileName"`
	AssemblyName AssemblyName `xml:"AssemblyName"`
	CacheFlash CacheFlash `xml:"CacheFlash"`
	Avrtoolinterfaceclock Avrtoolinterfaceclock `xml:"avrtoolinterfaceclock"`
	EraseKey EraseKey `xml:"EraseKey"`
	OutputType OutputType `xml:"OutputType"`
	Language Language `xml:"Language"`
	ToolchainFlavour ToolchainFlavour `xml:"ToolchainFlavour"`
	KeepTimersRunning KeepTimersRunning `xml:"KeepTimersRunning"`
	PreserveEEPROM PreserveEEPROM `xml:"preserveEEPROM"`
	Eraseonlaunchrule Eraseonlaunchrule `xml:"eraseonlaunchrule"`
	ResetRule ResetRule `xml:"ResetRule"`
	Avrdevice Avrdevice `xml:"avrdevice"`
	OutputDirectory OutputDirectory `xml:"OutputDirectory"`
	OverrideVtorValue OverrideVtorValue `xml:"OverrideVtorValue"`
	GpdscFilePath GpdscFilePath `xml:"GpdscFilePath"`
	Avrtoolinterface Avrtoolinterface `xml:"avrtoolinterface"`
	Name Name `xml:"Name"`
	OverrideVtor OverrideVtor `xml:"OverrideVtor"`
	AcmeProjectConfigs AcmeProjectConfigs `xml:"AcmeProjectConfigs"`
	Avrdeviceexpectedsignature Avrdeviceexpectedsignature `xml:"avrdeviceexpectedsignature"`
	Avrdeviceseries Avrdeviceseries `xml:"avrdeviceseries"`
	UncachedRange UncachedRange `xml:"UncachedRange"`
	ToolchainSettings ToolchainSettings `xml:"ToolchainSettings"`
}

type RootNamespace struct {
	Content string `xml:",chardata"`
}

type BootSegment struct {
	Content int64 `xml:",chardata"`
}

type AtStartFilePath struct {
	Content string `xml:",chardata"`
}

type AtzipPath struct {
	Content string `xml:",chardata"`
}

type Avrtool struct {
	Content string `xml:",chardata"`
}

type ProjectVersion struct {
	Content float64 `xml:",chardata"`
}

type ProgFlashFromRam struct {
	Content bool `xml:",chardata"`
}

type RamSnippetAddress struct {
	Content string `xml:",chardata"`
}

type AsfFrameworkConfig struct {
	FrameworkData FrameworkData `xml:"framework-data"`
}

type FrameworkData struct {
	Documentation Documentation `xml:"documentation"`
	OfflineDocumentation OfflineDocumentation `xml:"offline-documentation"`
	Dependencies Dependencies `xml:"dependencies"`
	Options Options `xml:"options"`
	Configurations Configurations `xml:"configurations"`
	Files Files `xml:"files"`
}

type Documentation struct {
	Help string `xml:"help,attr"`
}

type OfflineDocumentation struct {
	Help string `xml:"help,attr"`
}

type Dependencies struct {
	ContentExtension ContentExtension `xml:"content-extension"`
}

type ContentExtension struct {
	Eid string `xml:"eid,attr"`
	Uuidref string `xml:"uuidref,attr"`
	Version string `xml:"version,attr"`
}

type Options struct {
	Content string `xml:",chardata"`
}

type Configurations struct {
	Content string `xml:",chardata"`
}

type Files struct {
	Content string `xml:",chardata"`
}

type ComAtmelAvrdbgToolAtmelice struct {
	ToolOptions ToolOptions `xml:"ToolOptions"`
	ToolType ToolType `xml:"ToolType"`
	ToolNumber ToolNumber `xml:"ToolNumber"`
	ToolName ToolName `xml:"ToolName"`
}

type ToolOptions struct {
	InterfaceName InterfaceName `xml:"InterfaceName"`
	InterfaceProperties InterfaceProperties `xml:"InterfaceProperties"`
}

type InterfaceName struct {
	Content string `xml:",chardata"`
}

type InterfaceProperties struct {
	SwdClock SwdClock `xml:"SwdClock"`
}

type SwdClock struct {
	Content int64 `xml:",chardata"`
}

type ToolType struct {
	Content string `xml:",chardata"`
}

type ToolNumber struct {
	Content string `xml:",chardata"`
}

type ToolName struct {
	Content string `xml:",chardata"`
}

type ComAtmelAvrdbgToolJtagice3plus struct {
	ToolOptions ToolOptions `xml:"ToolOptions"`
	ToolType ToolType `xml:"ToolType"`
	ToolNumber ToolNumber `xml:"ToolNumber"`
	ToolName ToolName `xml:"ToolName"`
}

type ToolchainName struct {
	Content string `xml:",chardata"`
}

type ProjectGuid struct {
	Content string `xml:",chardata"`
}

type OutputFileExtension struct {
	Content string `xml:",chardata"`
}

type Avrtoolserialnumber struct {
	Content string `xml:",chardata"`
}

type SchemaVersion struct {
	Content float64 `xml:",chardata"`
}

type OutputFileName struct {
	Content string `xml:",chardata"`
}

type AssemblyName struct {
	Content string `xml:",chardata"`
}

type CacheFlash struct {
	Content bool `xml:",chardata"`
}

type Avrtoolinterfaceclock struct {
	Content int64 `xml:",chardata"`
}

type EraseKey struct {
	Content string `xml:",chardata"`
}

type OutputType struct {
	Content string `xml:",chardata"`
}

type Language struct {
	Content string `xml:",chardata"`
}

type ToolchainFlavour struct {
	Content string `xml:",chardata"`
}

type KeepTimersRunning struct {
	Content bool `xml:",chardata"`
}

type PreserveEEPROM struct {
	Content bool `xml:",chardata"`
}

type Eraseonlaunchrule struct {
	Content int64 `xml:",chardata"`
}

type ResetRule struct {
	Content int64 `xml:",chardata"`
}

type Avrdevice struct {
	Content string `xml:",chardata"`
}

type OutputDirectory struct {
	Content string `xml:",chardata"`
}

type OverrideVtorValue struct {
	Content string `xml:",chardata"`
}

type GpdscFilePath struct {
	Content string `xml:",chardata"`
}

type Avrtoolinterface struct {
	Content string `xml:",chardata"`
}

type Name struct {
	Content string `xml:",chardata"`
}

type OverrideVtor struct {
	Content bool `xml:",chardata"`
}

type AcmeProjectConfigs struct {
	AcmeProjectConfig AcmeProjectConfig `xml:"AcmeProjectConfig"`
}

type AcmeProjectConfig struct {
	TopLevelComponents TopLevelComponents `xml:"TopLevelComponents"`
	AcmeActionInfos AcmeActionInfos `xml:"AcmeActionInfos"`
	NonsecureFilesInfo NonsecureFilesInfo `xml:"NonsecureFilesInfo"`
}

type TopLevelComponents struct {
	AcmeProjectComponents []AcmeProjectComponent `xml:"AcmeProjectComponent"`
}

type AcmeProjectComponent struct {
	Cgroup string `xml:"Cgroup,attr"`
	CVersion string `xml:"CVersion,attr"`
	IsAutoGenerated bool `xml:"IsAutoGenerated,attr"`
	CClass string `xml:"CClass,attr"`
}

type AcmeActionInfos struct {
	AcmeProjectActionInfos []AcmeProjectActionInfo `xml:"AcmeProjectActionInfo"`
}

type AcmeProjectActionInfo struct {
	IsConfig bool `xml:"IsConfig,attr"`
	Hash string `xml:"Hash,attr"`
	Action string `xml:"Action,attr"`
	Source string `xml:"Source,attr"`
}

type NonsecureFilesInfo struct {
	Content string `xml:",chardata"`
}

type Avrdeviceexpectedsignature struct {
	Content string `xml:",chardata"`
}

type Avrdeviceseries struct {
	Content string `xml:",chardata"`
}

type UncachedRange struct {
	Content string `xml:",chardata"`
}

type ToolchainSettings struct {
	ArmGcc ArmGcc `xml:"ArmGcc"`
}

type ArmGcc struct {
	ArmgccCompilerDirectoriesIncludePaths ArmgccCompilerDirectoriesIncludePaths `xml:"armgcc.compiler.directories.IncludePaths"`
	ArmgccCompilerOptimizationPrepareFunctionsForGarbageCollection ArmgccCompilerOptimizationPrepareFunctionsForGarbageCollection `xml:"armgcc.compiler.optimization.PrepareFunctionsForGarbageCollection"`
	ArmgccCommonOutputfilesBin ArmgccCommonOutputfilesBin `xml:"armgcc.common.outputfiles.bin"`
	ArmgccCompilerSymbolsDefSymbols ArmgccCompilerSymbolsDefSymbols `xml:"armgcc.compiler.symbols.DefSymbols"`
	ArmgccCompilerWarningsAllWarnings ArmgccCompilerWarningsAllWarnings `xml:"armgcc.compiler.warnings.AllWarnings"`
	ArmgccLinkerLibrariesLibrarySearchPaths ArmgccLinkerLibrariesLibrarySearchPaths `xml:"armgcc.linker.libraries.LibrarySearchPaths"`
	ArmgccLinkerMiscellaneousLinkerFlags ArmgccLinkerMiscellaneousLinkerFlags `xml:"armgcc.linker.miscellaneous.LinkerFlags"`
	ArmgccAssemblerGeneralIncludePaths ArmgccAssemblerGeneralIncludePaths `xml:"armgcc.assembler.general.IncludePaths"`
	ArmgccPreprocessingassemblerGeneralIncludePaths ArmgccPreprocessingassemblerGeneralIncludePaths `xml:"armgcc.preprocessingassembler.general.IncludePaths"`
	ArmgccCommonOutputfilesSrec ArmgccCommonOutputfilesSrec `xml:"armgcc.common.outputfiles.srec"`
	ArmgccCompilerOptimizationLevel ArmgccCompilerOptimizationLevel `xml:"armgcc.compiler.optimization.level"`
	ArmgccCommonOutputfilesEep ArmgccCommonOutputfilesEep `xml:"armgcc.common.outputfiles.eep"`
	ArmgccLinkerLibrariesLibraries ArmgccLinkerLibrariesLibraries `xml:"armgcc.linker.libraries.Libraries"`
	ArmgccCommonOutputfilesHex ArmgccCommonOutputfilesHex `xml:"armgcc.common.outputfiles.hex"`
	ArmgccCommonOutputfilesLss ArmgccCommonOutputfilesLss `xml:"armgcc.common.outputfiles.lss"`
	ArmgccLinkerGeneralUseNewlibNano ArmgccLinkerGeneralUseNewlibNano `xml:"armgcc.linker.general.UseNewlibNano"`
	ArmgccLinkerOptimizationGarbageCollectUnusedSections ArmgccLinkerOptimizationGarbageCollectUnusedSections `xml:"armgcc.linker.optimization.GarbageCollectUnusedSections"`
	ArmgccLinkerMemorysettingsExternalRAM ArmgccLinkerMemorysettingsExternalRAM `xml:"armgcc.linker.memorysettings.ExternalRAM"`
	ArmgccAssemblerDebuggingDebugLevel ArmgccAssemblerDebuggingDebugLevel `xml:"armgcc.assembler.debugging.DebugLevel"`
	ArmgccCompilerOptimizationDebugLevel ArmgccCompilerOptimizationDebugLevel `xml:"armgcc.compiler.optimization.DebugLevel"`
	ArmgccCompilerGeneralChangeDefaultCharTypeUnsigned ArmgccCompilerGeneralChangeDefaultCharTypeUnsigned `xml:"armgcc.compiler.general.ChangeDefaultCharTypeUnsigned"`
	ArmgccPreprocessingassemblerDebuggingDebugLevel ArmgccPreprocessingassemblerDebuggingDebugLevel `xml:"armgcc.preprocessingassembler.debugging.DebugLevel"`
	ArmgccCompilerGeneralChangeDefaultBitFieldUnsigned ArmgccCompilerGeneralChangeDefaultBitFieldUnsigned `xml:"armgcc.compiler.general.ChangeDefaultBitFieldUnsigned"`
	ArmgccCompilerOptimizationPrepareDataForGarbageCollection ArmgccCompilerOptimizationPrepareDataForGarbageCollection `xml:"armgcc.compiler.optimization.PrepareDataForGarbageCollection"`
}

type ArmgccCompilerDirectoriesIncludePaths struct {
	ListValues ListValues `xml:"ListValues"`
}

type ListValues struct {
	Values []Value `xml:"Value"`
}

type Value struct {
	Content string `xml:",chardata"`
}

type ArmgccCompilerOptimizationPrepareFunctionsForGarbageCollection struct {
	Content bool `xml:",chardata"`
}

type ArmgccCommonOutputfilesBin struct {
	Content bool `xml:",chardata"`
}

type ArmgccCompilerSymbolsDefSymbols struct {
	ListValues ListValues `xml:"ListValues"`
}

type ArmgccCompilerWarningsAllWarnings struct {
	Content bool `xml:",chardata"`
}

type ArmgccLinkerLibrariesLibrarySearchPaths struct {
	ListValues ListValues `xml:"ListValues"`
}

type ArmgccLinkerMiscellaneousLinkerFlags struct {
	Content string `xml:",chardata"`
}

type ArmgccAssemblerGeneralIncludePaths struct {
	ListValues ListValues `xml:"ListValues"`
}

type ArmgccPreprocessingassemblerGeneralIncludePaths struct {
	ListValues ListValues `xml:"ListValues"`
}

type ArmgccCommonOutputfilesSrec struct {
	Content bool `xml:",chardata"`
}

type ArmgccCompilerOptimizationLevel struct {
	Content string `xml:",chardata"`
}

type ArmgccCommonOutputfilesEep struct {
	Content bool `xml:",chardata"`
}

type ArmgccLinkerLibrariesLibraries struct {
	ListValues ListValues `xml:"ListValues"`
}

type ArmgccCommonOutputfilesHex struct {
	Content bool `xml:",chardata"`
}

type ArmgccCommonOutputfilesLss struct {
	Content bool `xml:",chardata"`
}

type ArmgccLinkerGeneralUseNewlibNano struct {
	Content bool `xml:",chardata"`
}

type ArmgccLinkerOptimizationGarbageCollectUnusedSections struct {
	Content bool `xml:",chardata"`
}

type ArmgccLinkerMemorysettingsExternalRAM struct {
	Content string `xml:",chardata"`
}

type ArmgccAssemblerDebuggingDebugLevel struct {
	Content string `xml:",chardata"`
}

type ArmgccCompilerOptimizationDebugLevel struct {
	Content string `xml:",chardata"`
}

type ArmgccCompilerGeneralChangeDefaultCharTypeUnsigned struct {
	Content bool `xml:",chardata"`
}

type ArmgccPreprocessingassemblerDebuggingDebugLevel struct {
	Content string `xml:",chardata"`
}

type ArmgccCompilerGeneralChangeDefaultBitFieldUnsigned struct {
	Content bool `xml:",chardata"`
}

type ArmgccCompilerOptimizationPrepareDataForGarbageCollection struct {
	Content bool `xml:",chardata"`
}

type ItemGroup struct {
	Compiles []Compile `xml:"Compile"`
	Folders []Folder `xml:"Folder"`
	Nones []None `xml:"None"`
}

type Compile struct {
	Include string `xml:"Include,attr"`
	SubType SubType `xml:"SubType"`
}

type SubType struct {
	Content string `xml:",chardata"`
}

type Folder struct {
	Include string `xml:"Include,attr"`
}

type None struct {
	Include string `xml:"Include,attr"`
	SubType SubType `xml:"SubType"`
}

type Import struct {
	Project string `xml:"Project,attr"`
}

