package provide base64 2.4.2
package require Tcl 8.2
if {![catch {package require Trf 2.0}]} {}
if {![package vsatisfies [package provide Tcl] 8.2]} {return}
package ifneeded base64   2.4.2 [list source [file join $dir base64.tcl]]
package ifneeded uuencode 1.1.5 [list source [file join $dir uuencode.tcl]]
package ifneeded yencode  1.1.3 [list source [file join $dir yencode.tcl]]
package ifneeded ascii85  1.0   [list source [file join $dir ascii85.tcl]]
package require Tcl 8.2;                # tcl minimum version
catch {package require crc32};          # tcllib 1.1
catch {package require tcllibc};        # critcl enhancements for tcllib
if {[package provide critcl] != {}} {}
package provide yencode 1.1.3
if {![package vsatisfies [package provide Tcl] 8.5]} {return}

package ifneeded yaml         0.4.1 [list source [file join $dir yaml.tcl]]
package ifneeded huddle       0.3   [list source [file join $dir huddle.tcl]]
package ifneeded huddle::json 0.1   [list source [file join $dir json2huddle.tcl]]

