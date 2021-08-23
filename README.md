# as2make - Atmel Studio to Makefile

**Disclaimer:** This project is not affiliated to Microchip in any way.

Atmel Studio, built on Visual Studio, is a Windows-only solution. A lot of companies, however, have their buildservers (like Hudson, Jenkins, Bamboo, ...) hosted on Linux environments. While Atmel Studio can generate Makefiles which enable building on Linux, keeping them up-to-date is a tedious job.

Since we had the same problem we build _as2make_. The tool reads the actual Atmel Studio project files and SDK files (at the moment it supports CMSIS and DFPs) and generated the Makefile (and makedep.mk) on-the-fly in Linux. Right in the build itself.

Example Call:
```plain
> as2make.linux \
  --output-name=firmware \
  --toolchain=../toolchain \
  --toolchain-prefix=arm-none-eabi- \
  --cmsis-sdk=../cmsis-sdk/CMSIS \
  --dfp-sdk=../dfp-sdk --project=./firmware.cproj

Loading project file: ./firmware.cproj ...
Loading DPF package descriptor: ../dfp-sdk/Atmel.SAMD21_DFP.pdsc ...
Generating Makefile.gen ...
Generating makedep.mk ...
```

At the moment the feature list is small and the tool supports exactly the stuff we needed to build our own firmware files. However, we are happy to add additional features necessary to provide a full featured solution. We're also happy to see Pull Requests and Feature Requests. For PRs and FRs, it'd be great to provide examples of the project files and Visual Studio generated Makefiles as a hint what it should look like after generation.

Anyway, we hope that this tool is of help for others, too.

## Information

Information on how to find the necessary SDKs:

https://ww1.microchip.com/downloads/en/DeviceDoc/Using-device-packs-with-toolchain-4.9.3.26.txt

Device Support Packs (DFP):

http://packs.download.atmel.com/

Rename *.atpack to *.zip and unpack.

CMSIS SDK:

http://www.keil.com/dd2/pack/

Rename *.pack to *.zip and unpack.

GCC ARM Toolchain:

https://developer.arm.com/tools-and-software/open-source-software/developer-tools/gnu-toolchain/gnu-rm/downloads

(GCC 6 recommended)
