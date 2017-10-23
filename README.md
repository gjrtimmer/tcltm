# tcltm - Tcl Module Builder

[![pipeline status](https://gitlab.timmertech.nl/VANAD/cpm/tcltm/badges/master/pipeline.svg)](https://gitlab.timmertech.nl/VANAD/cpm/tcltm/commits/master)
[![](https://images.microbadger.com/badges/image/datacore/tcltm.svg)](https://microbadger.com/images/datacore/tcltm)
[![](https://images.microbadger.com/badges/version/datacore/tcltm.svg)](https://microbadger.com/images/datacore/tcltm)
[![](https://images.microbadger.com/badges/commit/datacore/tcltm.svg)](https://microbadger.com/images/datacore/tcltm)
[![](https://images.microbadger.com/badges/license/datacore/tcltm.svg)](https://microbadger.com/images/datacore/tcltm)

# Table of Contents
- [Dependencies](#dependencies)
- [Usage](#usage)
- [Arguments](#arguments)
- [Tcl Module Specification](/TM.md)

# Dependencies
- yaml 0.3.6+ (tcllib)

# Usage

```bash
$ tcltm -d . -o ../
```

# Arguments

| Short | Long | Description | Usage |
|-------|------|-------------|-------|
| ```-d``` DIR | ```--dir``` DIR | Input Directory | ```-d ./tcllib1.18/mime``` |
| ```-o``` DIR | ```--out``` DIR | Output directory | ```-o ./tcl8/8.6```
| ```-c``` | ```--create-dirs``` | Create target directory structure (Example: tcl/8.6) | |
| ```-s``` FILE | ```--scan``` FILE | Scan for required packages | 

# .tcltm
A Tcl Module is created based upon the .tcltm file present within the directory.

A ```.tcltm``` configuration file contains the entire specification for a package. Within a directory multiple packages can be defined in Tcl. ```.tcltm``` supports the specification of multiple packages in a single file.

For full specification see [```.tcltm specification```](/TCLTM.md)
