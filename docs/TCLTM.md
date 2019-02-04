# tcltm Specification

## Table of Contents

- [Required Keys](#required-keys)
- [Configuration Keys](#optional-keys)
- [Binary Files](#binary-files)
- [Example](#example)

A `tcltm` specification is defined in [Yaml](https://en.wikipedia.org/wiki/YAML).

Multiple packages can be defined within a single `.tcltm` file, the `Package` key
is an array of Tcl packages.

## Required Keys

The following keys are required.

- name
- tcl
- files

## Configuration Keys

| Key          | Required | Description                                                                                                            |
| ------------ | -------- | ---------------------------------------------------------------------------------------------------------------------- |
| name         | x        | Name of package                                                                                                        |
| version      | x        | Version of package                                                                                                     |
| tcl          | x        | Tcl version                                                                                                            |
| summary      |          | Summary description                                                                                                    |
| description  |          | Full description                                                                                                       |
| license      |          | License, if not present in configuration `tcltm` will look for a `LICENSE` file                                        |
| dependencies |          | Dependencies, version of a dependency can be specified after the name with a space.                                    |
| meta         |          | Additional meta keys to be included in the header                                                                      |
| extension    |          | Extension of generated module, defaults to `tm`                                                                        |
| finalname    |          | Filename of generated output. Defaults to `{Name}-{Version}.{Extension}`                                               |
| filter       |          | Global custom filter keys, applies to all files which have filtering turned on. See [File Filtering](#file-filtering). |
| files        | x        | Source files to include. See [File Configuration](#file-configuration).                                                |
| bootstrap    |          | Initscript runs before source, either embedded or the name of a `.tcl` file.                                           |
| init         |          | Initscript run after **source**, either embedded or the name of a `.tcl` file.                                         |

### Version

Because multiple packages can be build from the same source (`.tcltm` configuration), the key `version` is a special case. In some cases a CI environment
will be used to build a package, which means that the actual version might come from an environment key or from a git tag for example.

The key `version` support environment variable loading, meaning that a version value can be read from the environment.
This can be achieved in the same way as being used by the [File Filtering - Environment](#file-filtering---environment).

To load a version value from the environment use the following as value for the version `env:VARIABLE_NAME`.

If you for example are building the package with a integrated CI from `gitlab`, and want to use the gitlab runner variable called `CI_COMMIT_REF_NAME` which holds the branch name
or tag name during build, you define the `version` tag in the `.tcltm` configuration as below.

```yaml
package:
  - name: testpkg
    version: env:CI_COMMIT_REF_NAME
    tcl: 8.5
    summary: test package on Gitlab CI
```

Say for example that you are building a tag with the value `3.0.0`. Then the resulting filename of the package being build would be `testpkg-3.0.0.tm`.

A default value as a fallback can be provided by adding an additional `:` with a default value.

Example: `env:CI_COMMIT_REF_NAME:0.0.0`

The example below will use the value of the environment variable `CI_COMMIT_REF_NAME`,
if the variable is not present then the version will be set to `0.0.0`.

The package version can also be loaded from a pkgIndex.tcl file in the source folder.
This feature can be used by providing the commandline option `--version-from-index`.

**Note: When using `--version-from-index` the package name in the `.tcltm` configuration has to be the same as the package name in the `pkgIndex.tcl`**

### License

The license can be included in the `.tcltm` file in the `License` key by using the multiline indicator `|-`, or specifying a file.
If not configured `tcltm` will look for a file named `LICENSE`.

### File Configuration

| Key       | Required | Description                                                                      | Link                              |
| --------- | -------- | -------------------------------------------------------------------------------- | --------------------------------- |
| name      | X        | Name of file to include                                                          |                                   |
| type      |          | Type of file. Default to `source`                                                | [File Type](#file-type)           |
| action    |          | Action to perform for the included file. Defaults to `none`                      | [File Action](#file-action)       |
| target    |          | File target.                                                                     |                                   |
| filtering |          | File filtering will be applied before the file is embedded. Defaults to `false`. | [File Filtering](#file-filtering) |
| filter    |          | Provide custom filter keys.                                                      | [File Filtering](#file-filtering) |

### Filtering

Accepted values for key `filtering`

- 0
- false
- off
- 1
- true
- on

### File Type

- binary
- source

### File Action

| Action  | Status              | Description                                                                                                            |
| ------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| none    | READY               | No action to perform for included file, this means simply include the code. This is the default.                       |
| run     | Not Yet Implemented | See [File Action Run](#file-action-run)                                                                                |
| load    | READY               | This will cause the package to load the file, with the `tcl` command `load`. `tcltm` will generate the necessary code. |
| extract | Not Yet Implemented | This will extract the included file.                                                                                   |

### File Action Run

This will cause the file to be included not as source code.

In the final target auto generated code will be included which will extract this file from the target and then run the included file. The run action will be done by evaluating the included target file. This action will be performed before the `InitScript` of the package if defined.

### File Filtering

Before a file is embedded filtering can be applied. Filtering consists of replacing markers in the source file with the file content.

Default provided filter keys.

| Filter Key | Description                                              |
| ---------- | -------------------------------------------------------- |
| @PNAME@    | Will be replaced with the package name.                  |
| @PVERSION@ | Will be replaced with the package version.               |
| @FILENAME@ | Will be replaced with the filename of the included file. |

With the configuration key `Filter`, custom filter keys can be provided. Any sub key of `Filter` will be converted to `@` + `Key Name` + `@` == `@KEY@`.

`@KEY@` can be used in the source code and will be replaced with the provided value.

> Note: The name of the filter key is case-sensitive

Example:

```yaml
files:
  - name: test.tcl
    filtering: true
    filter:
      FOO: bar
```

The provided example will replace `@FOO@` in the file content with `bar`.

#### File Filtering - ENVIRONMENT

File filtering supports using environment variables. This can be used be setting a key in the filter which you want to use
for replacement in your code, like normal filtering. But as the value of the key you will enter `env:VARIABLE_NAME`.

In the example below, the environment variable `USER` (Linux: `$USER`; Windows `%USER%`) is used as a value for the filter key `USER`.
In the source code you will use `@USER@`.

Example:

```yaml
files:
  - name: test.tcl
    filtering: true
    filter:
      USER: env:USER
```

This example will result that every occurance of the key `@USER@` in the source code of `test.tcl` is replaced with the value of the
enviroment variable `USER`.

When using the environment to provide a value, its also possible to provide a default value.

A default value can be provided by adding and additional `:` and a default value.

Example: `env:USER:username`

If the variable `USER` can not be found in the environment the value will be set to `username`.

### Ignore Code Parts

In several cases it my be required to ignore several parts from the source to be included in the Tcl module.
This can be achieved by placing the code to be ignore between the following comments.

#### Ignore Single line

A single line can be ignored for inclusion in two different ways.
A comment with `TCLTM IGNORE NEXT` can be placed on top the the line with the instruction to ignore the next line.

```tcl
# TCLTM IGNORE NEXT
package provide pkg 1.0.0
```

Its also possible to provide an inline comment to ignore the current line. This can be done with
the comment `TCLTM IGNORE`.

```tcl
return ; # TCLTM IGNORE
```

#### Ignore Multiple Lines

Ignoring multiple lines can be done by placing `TCLTM IGNORE BEGIN` and `TCLTM IGNORE END` comment markers around
the code block to be ignored.

```tcl
# TCLTM IGNORE BEGIN
{code to be ignored for inclusion}
# TCLTM IGNORE END
```

## Specification

```yaml
---
package:
  - name: Package name
    version: Package version
    tcl: Tcl version
    summary: Summary of package
    description: Full package description
    license: License
    dependencies: List of dependencies this package depends on.
    meta: Additional meta keys for package header
    extension: Extension of output
    finalname: Output name
    filter: Global filter keys for all files
    bootstrap: -|
        Initscript to run BEFORE source, can be multiline within the config
        or can be simply a script file which has a .tcl extension.
    init: -|
        Initscript to run AFTER source, can be multiline within the config
        or can be simply a script file which has a .tcl extension.
    files:
      - name:
        type: script, binary
        action: none, run, load, extract
        target: ~
        filtering: boolean (0, false, off, 1 true, on)
        filter:
          KEY: VALUE
```

## Binary Files

All files defined with the key `files` which do not have the extension `.tcl` are treated as
binary files.

Additionally there is the possibility to auto-load libraries.
This can be achieved be setting the `action` of the file to `load`.

The `.tm` file will extract the embedded libary and directly issue the Tcl `load` command.

If there is no command given for a binary, the `tcltm` builder will generate code
which will add the extract binary path to the variable `::tcltm::binary::path` which the user then
can use within the init script to handle the binary themself.

> Binary files are included in the order in which they are defined in the configuration.
