# tcltm config

[![GoDoc](https://godoc.org/gitlab.timmertech.nl/tcl/tcltm?status.svg)](https://godoc.org/gitlab.timmertech.nl/tcl/tcltm)

## Configuration

Configuration filename: `.tcltm`

A configuration file is split into 3 parts.

* version
* tcltm
* modules

### version

`version` defines the version of the configuration file, intended for future updates and evolution of the configuration.

### tcltm

The `tcltm` section defines the global configuration properties of the Tcl Module Builder.

**These values can be overriden for each module**

#### include

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| require | `bool` | `false` | When `true` this will ensure that the `package require` commands will be preserved from the included resources. |

#### exclude

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| comments | `bool` | `false` | When `true` comments are stripped from the included resources |
| resource-prefix | `string` | `""` | This will strip the provided prefix from the resource path in the output |
| resource-path | `bool` | `false` | This will strip the complete path prefix from the resource path in the output |
| dependencies | `bool` | `false` | If `true` this will cause the defined dependenices not to be written with a `package require` command |
| provide | `bool` | `false` | If `true` this will not write the default `package provide` command for the genereated output. |
| satisfy-tcl | `bool` | `false` | If `true` the default `package vsatisfies` command is not written to the output |

#### output

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| repository | `bool` | `false` | Repository will when enabled write the modules defined within the configuration to an output directory conforming to the default Tcl TM Repository format |
| interactive-loader | `bool` | `false` | This will reconfigure the binary loader to only be executed when the module is loaded within an interactive shell |

#### Example

```yaml
tcltm:
  include:
    require: true
  exclude:
    comments: true
    resource-prefix: true
    dependencies: true
    provide: true
    satisfy-tcl: true
  output:
    repository: true
    interactive-loader: true
```

## Example

```yaml
---
version: 4.0

tcltm:
  include:
    require: true
  exclude:
    comments: true
    resource-prefix: true
    dependencies: true
    provide: true
    satisfy-tcl: true
  output:
    repository: true
    interactive-loader: true

modules:
  - name: config-test
    version: 0.0.0
    tcl: 8.6
    authors:
      - Name: John Doe
        Email: john.doe@example.com
    interp: tclsh
    summary: Config Test
    description: |-
      Multiline description
      for config test
    license: LICENSE
    dependencies:
      - Tclx:8.4
      - base64
    extension: tm
    finalname: config-test.tm
    meta:
      foo: bar
    filter:
      user: test

    # Override global tcltm properties
    include:
      require: true
    exclude:
      comments: true
      resource-prefix: true
      dependencies: true
      provide: true
      satisfy-tcl: true
    output:
      repository: true
      interactive-loader: true

    # Resource Configuration
    files:
      - name: test.tcl
        type: binary
        action: run
        filtering: true
        filter:
          os: linux
```
