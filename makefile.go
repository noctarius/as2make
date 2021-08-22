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
	"os"
	"strings"
)

func writeMakefile(build builder.Build) {
	fi, _ := os.Stat("Makefile.gen")
	if fi != nil {
		os.Remove("Makefile.gen")
	}

	file, err := os.OpenFile("Makefile.gen", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write(file, "ADDITIONAL_DEPENDENCIES :=\n")
	write(file, "EXECUTABLES :=\n")
	write(file, "LIB_DEP :=\n")
	write(file, "USER_OBJS :=\n")
	write(file, "OUTPUT_FILE_PATH :=\n")
	write(file, "OUTPUT_FILE_PATH +=build.elf\n")
	write(file, "OUTPUT_FILE_PATH_AS_ARGS +=build.elf\n")
	write(file, "OUTPUT_FILE_DEP:= ./makedep.mk\n\n")

	writeSubdirsFiles(build, file)
	writeSrcFiles(build, file)
	writeObjFiles(build, file)
	writeDepFiles(build, file)
	writeAsFiles(build, file)
	writeLdFiles(build, file)
	writeCompileSteps(build, file)
	writeClean(file)
	writeAllTarget(file)
	writeCompileTarget(build, file)
	writeCleanTarget(build, file)
}

func writeCompileSteps(build builder.Build, file *os.File) {
	compileCommand := buildCommand(build)
	sources := build.Sources()
	if len(sources) > 0 {
		for _, source := range sources {
			sourceFile := source.SourceFile()
			objectFile := source.ObjectFile()
			write(file, "%s: %s\n", objectFile, sourceFile)
			write(file, "\t@echo Building file: $<\n")
			write(file, "\t@%s\n", compileCommand)
			write(file, "\t@echo Finished building: $<\n")
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeSubdirsFiles(build builder.Build, file *os.File) {
	subdirs := build.Subdirs()
	if len(subdirs) > 0 {
		write(file, "SUBDIRS := \\\n")
		for i, subdir := range build.Subdirs() {
			write(file, "%s", subdir)
			if i < len(subdirs)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeSrcFiles(build builder.Build, file *os.File) {
	sources := build.Sources()
	if len(sources) > 0 {
		write(file, "C_SRCS += \\\n")
		for i, source := range sources {
			write(file, "%s", source.SourceFile())
			if i < len(sources)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeObjFiles(build builder.Build, file *os.File) {
	sources := build.Sources()
	if len(sources) > 0 {
		write(file, "OBJS += \\\n")
		for i, source := range sources {
			write(file, "%s", source.ObjectFile())
			if i < len(sources)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeDepFiles(build builder.Build, file *os.File) {
	sources := build.Sources()
	if len(sources) > 0 {
		write(file, "C_DEPS_AS_ARGS += \\\n")
		for i, source := range sources {
			write(file, "%s", source.DependFile())
			if i < len(sources)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeAsFiles(build builder.Build, file *os.File) {
	sources := build.Sources()
	if len(sources) > 0 {
		write(file, "OBJS_AS_ARGS += \\\n")
		for i, source := range sources {
			write(file, "%s", source.ObjectFile())
			if i < len(sources)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeLdFiles(build builder.Build, file *os.File) {
	sources := build.LinkerSources()
	if len(sources) > 0 {
		write(file, "LINKER_SCRIPT_DEP += \\\n")
		for i, source := range sources {
			write(file, "%s", source.SourceFile())
			if i < len(sources)-1 {
				write(file, " \\")
			}
			write(file, "\n")
		}
		write(file, "\n")
	}
}

func writeClean(file *os.File) {
	if _, err := file.WriteString("ifneq ($(MAKECMDGOALS),clean)\n"); err != nil {
		panic(err)
	}
	if _, err := file.WriteString("ifneq ($(strip $(C_DEPS)),)\n"); err != nil {
		panic(err)
	}
	if _, err := file.WriteString("-include $(C_DEPS)\n"); err != nil {
		panic(err)
	}
	if _, err := file.WriteString("endif\n"); err != nil {
		panic(err)
	}
	if _, err := file.WriteString("endif\n\n"); err != nil {
		panic(err)
	}
}

func writeAllTarget(file *os.File) {
	if _, err := file.WriteString("all: $(OUTPUT_FILE_PATH) $(ADDITIONAL_DEPENDENCIES)\n\n"); err != nil {
		panic(err)
	}
}

func writeCompileTarget(build builder.Build, file *os.File) {
	gcc := build.Toolchain().Executable("gcc")
	objCopy := build.Toolchain().Executable("objcopy")
	objDump := build.Toolchain().Executable("objdump")
	size := build.Toolchain().Executable("size")

	write(file, "$(OUTPUT_FILE_PATH): $(OBJS) $(USER_OBJS) $(OUTPUT_FILE_DEP) $(LIB_DEP) $(LINKER_SCRIPT_DEP)\n")
	write(file, "\t@echo Building target: $@\n")
	write(file, "\t%s", gcc)
	write(file, " -o$(OUTPUT_FILE_PATH_AS_ARGS) $(OBJS_AS_ARGS) $(USER_OBJS) $(LIBS) -mthumb -Wl,-Map=\"build.map\" --specs=nano.specs -Wl,--start-group -lm  -Wl,--end-group -L\"./Device_Startup\"  -Wl,--gc-sections -mcpu=cortex-m0plus -Tsamd21g18a_flash.ld\n")
	if build.WithHex() {
		write(file, "\t%s", objCopy)
		write(file, " -O ihex -R .eeprom -R .fuse -R .lock -R .signature  \"build.elf\" \"build.hex\"\n")
	}
	if build.WithLss() {
		write(file, "\t%s", objDump)
		write(file, " -h -S \"build.elf\" > \"build.lss\"\n")
	}
	if build.WithEep() {
		write(file, "\t%s", objCopy)
		write(file, " -j .eeprom --set-section-flags=.eeprom=alloc,load --change-section-lma .eeprom=0 --no-change-warnings -O binary \"build.elf\" \"build.eep\" || exit 0\n")
	}
	if build.WithBin() {
		write(file, "\t%s", objCopy)
		write(file, " -O binary \"build.elf\" \"build.bin\"\n")
	}
	if build.WithSrec() {
		write(file, "\t%s", objCopy)
		write(file, " -O srec -R .eeprom -R .fuse -R .lock -R .signature  \"build.elf\" \"build.srec\"\n")
	}
	write(file, "\t%s", size)
	write(file, " \"build.elf\"\n")
	write(file, "\t@echo Finished successfully: $@\n\n")
}

func writeCleanTarget(build builder.Build, file *os.File) {
	write(file, "clean:\n")
	if _, err := file.WriteString("\t-rm -rf $(OBJS_AS_ARGS) $(EXECUTABLES)\n"); err != nil {
		panic(err)
	}
	if _, err := file.WriteString("\t-rm -rf $(C_DEPS_AS_ARGS)\n"); err != nil {
		panic(err)
	}
	if build.WithHex() {
		if _, err := file.WriteString("\t-rm -rf build.hex\n"); err != nil {
			panic(err)
		}
	}
	if build.WithLss() {
		if _, err := file.WriteString("\t-rm -rf build.lss\n"); err != nil {
			panic(err)
		}
	}
	if build.WithEep() {
		if _, err := file.WriteString("\t-rm -rf build.eep\n"); err != nil {
			panic(err)
		}
	}
	if build.WithBin() {
		if _, err := file.WriteString("\t-rm -rf build.bin\n"); err != nil {
			panic(err)
		}
	}
	if build.WithSrec() {
		if _, err := file.WriteString("\t-rm -rf build.srec\n"); err != nil {
			panic(err)
		}
	}
	if _, err := file.WriteString("\t-rm -rf build.elf build.a build.map\n"); err != nil {
		panic(err)
	}
}

func buildCommand(build builder.Build) string {
	toolchain := build.Toolchain()
	includes := build.IncludePaths()
	symbols := build.DefSymbols()
	opLevel := build.OptimizationLevel()
	return fmt.Sprintf(
		"%s -x c -mthumb -D__SAMD21G18A__ %s %s %s -ffunction-sections -mlong-calls -Wall -mcpu=cortex-m0plus -c -std=gnu99 -MD -MP -MF \"$(@:%%.o=%%.d)\" -MT\"$(@:%%.o=%%.d)\" -MT\"$(@:%%.o=%%.o)\" -o \"$@\" \"$<\" ",
		toolchain.Executable("gcc"),
		strings.Join(symbols, " "),
		strings.Join(includes, " "),
		opLevel,
	)
}
