# .tcltm Specification

A ```tcltm``` specification is defined in [Yaml](https://en.wikipedia.org/wiki/YAML).

# Example

```yaml
--- # Package SMTP
Package:
  - 
    Name: smtp
    Version: 1.4.5
    Summary: ~
    Description: ~
    License: ~
    Tcl: 8.3
    Dependencies:
      - mime 1.4.1
      - SASL 1.0
      - SALS::NTLM 1.0
    Files:
      - smtp.tcl
    InitScript: ~
```
