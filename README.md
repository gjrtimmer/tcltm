# tcltm - Tcl Module Builder

[![pipeline status](https://gitlab.timmertech.nl/tcl/tcltm/badges/master/pipeline.svg)](https://gitlab.timmertech.nl/tcl/tcltm/commits/master)

`tcltm` is a Tcl Module builder. It will create a Tcl module `*.tm` file based upon the `.tcltm` configuration file present within a project.
The module file which is generated is a pure Tcl implementation.

Please note that if a module is created which has dependencies that the create module will depend on these dependencies.

## Table of Contents

- [tcltm - Tcl Module Builder](#tcltm---tcl-module-builder)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
    - [Runtime](#runtime)
    - [Development Libraries](#development-libraries)
  - [Usage](#usage)
    - [Arguments](#arguments)
    - [Environment Variables](#environment-variables)
    - [Docker](#docker)
  - [.tcltm Specification](#tcltm-specification)
  - [Development](#development)

## Requirements

These dependencies are required to use `tcltm` to create a Tcl Module.

- Tcl \>= 8.5
- tcllib
  - sha1
  - try
  - yaml >= 0.3.6

### Runtime

These dependencies are required by the target system to load or run the created Tcl Module.

- Tcl \> 8.5

### Development Libraries

- tcltest

## Usage

```bash
tcltm -d . -o ../
```

### Arguments

See help `tcltm --help`

### Environment Variables

The following configuration options can be given by environment variables.
For example the version number of the package to be build can be set to the git tag by using this option.

Available overrides from environment variables.

| Variable | Override         |
| -------- | ---------------- |
| VERSION  | Package Version. |

### Docker

`tcltm` can be used from a docker container. The Tcl project is to be mounted as the `/data` volume.
The docker image which is built from this repository, is available as `datacore/tcltm`.

[Docker Hub](https://cloud.docker.com/repository/docker/datacore/tcltm)

The example below the current directory on the host contains the Tcl project with the `.tcltm` config.

Example:

```bash
docker pull datacore/tcltm
docker run --rm --name tcltm -v $(pwd):/data -it datacore/tcltm tcltm
```

## .tcltm Specification

A Tcl Module is created based upon the .tcltm file present within the directory.

A `.tcltm` configuration file contains the entire specification for a package. Within a directory multiple packages can be defined in Tcl. `.tcltm` supports the specification of multiple packages in a single file.

For full specification see [`.tcltm specification`](/docs/TCLTM.md)

## Development

See the [Development Guide](./docs/DEVELOPMENT.md)
