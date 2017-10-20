# Tcl Module File Specification

# File Layout
<!-- language: lang-none -->

    -------------------
    |     LICENSE     |
    -------------------
    |     HEADER      |
    -------------------
    |      INIT       |
    -------------------
    |     SOURCE      |
    -------------------
    | CTRL-Z (\u001A) |
    -------------------
    |   BINARY DATA   |
    -------------------

# License
When using ```tcltm``` to build a Tcl Module it will scan the provided directory. If a LICENSE file is found it will be included as the LICENSE for the module.

***NOTE*** ```tcltm``` will remove all the first lines of comments from the Tcl source files.

# Header

# Init

# Source

# Binary Data
