NAME:
    tcltm - Tcl Module Builder

USAGE:
    tcltm ?options?

VERSION:
    $::tcltm::version ($::tcltm::commit)

OPTIONS:
    -i DIR, --in DIR        Input directory. (Defaults: current directory)
    -o DIR, --out DIR       Output directory. (Defaults: current directory)
    -c FILE, --config FILE  Alternate config file. (Defaults: .tcltm)
    -p NAME, --pkg NAME     Only build package <NAME> from config.
                            (Defaults: build all)

    --version-from-index    Use package version from pkgIndex.tcl
                            Only works when package name between config and
                            pkgIndex.tcl is the same.
    --strip-comments        Strip comments from source
    --exclude-satisfy-tcl   Exclude Tcl vsatisfies command
    --exclude-deps          Exclude package require commands for dependencies
    --exclude-provide       Exclude package provide command

    --repository            Create repository output directories.
                            (tcl8/tcl<version>/module.tm)

    --scan FILE             Scan FILE for Tcl dependencies.
                            If file is a directory, all .tcl files in the
                            directory will be scanned.

    --verbose               Verbose logging
    --version               Show version
    -h, --help              Show help
