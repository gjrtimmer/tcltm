namespace eval ::tcltm::loader {
    variable script {
namespace eval ::tcltm::binary {
    variable path
    variable resources
    variable name

    proc loader {} {
        variable path
        variable resources [list]
        variable name

        # Check for ::tcltm::binary::path override from bootstrap
        if { ![info exists path] || [string length $path] == 0 } {
            set path [file normalize [file dirname [info script]]]
        }

        # Read header
        set bin [open [info script] {RDONLY BINARY}]
        set header 0
        while { [gets $bin line] >= 0 } {
            if { [string match {*TCLTM*HEADER*BEGIN*} $line] } {
                set header 1
                continue
            }
            if { [string match {*TCLTM*HEADER*END*} $line] } {
                break
            }
            if { [string match {*NAME*} $line] } {
                regexp {^# ([[:alpha:]]+): ([[:alpha:]]+$)} $line -> - name
            }
            if { [string match {*RESOURCE*} $line] } {
                set res {*}[string trimleft [lindex [split $line ":"] 1]]
                dict lappend resources files [dict get $res NAME]
                dict set resources [dict get $res NAME] $res
            }
        }

        # Reset Binary Index
        seek $bin 0

        # Read entire file
        set bindata [read $bin]
        close $bin

        # Binary Index
        set bindex [string first \\u001A $bindata]
        incr bindex

        # Extract all resouces
        foreach f [dict get $resources files] {
            set finfo [dict get $resources $f]
            set tmp [file normalize [file join $path [dict get $finfo NAME]]]
            set fh [open $tmp w]
            fconfigure $fh -translation binary
            fconfigure $fh -encoding binary
            puts -nonewline $fh [string range $bindata $bindex [incr bindex [dict get $finfo SIZE]]-1]
            flush $fh
            close $fh

            # Verify resource
            if { [package vsatisfies [package require sha1] 2.0.3] } {
                set hash [::sha1::sha1 -hex -file $tmp]
                if { $hash ne [dict get $finfo HASH] } {
                    return -code error "[file tail [info script]]: Hash invalid for embedded binary [dict get $finfo NAME]"
                }
            }

            # Preform Action
            if { [dict exists $finfo ACTION] } {
                switch -exact -- [string toupper [dict get $finfo ACTION]] {
                    NONE {
                        # No Action
                    }
                    RUN {
                        # Action RUN
                        if { [catch {source $tmp} err] } {
                            return -code error "Failed to run embedded resource: $tmp"
                        }
                    }
                    LOAD {
                        # Try normal load first
                        if { [catch {load $tmp}] } {
                            # Load failed
                            # Retry load with package name
                            if { [catch {load $tmp $name}] } {
                                return -code error "[file tail [info script]]: failed to load embedded binary [dict get $finfo NAME]"
                            }
                        }
                    }
                    EXTRACT {
                        # TODO: Implement Action EXTRACT
                    }
                    default {
                        # No Action
                    }
                }
            }
            incr bindex
        }
    }
}
::tcltm::binary::loader
} ; # END Variable script

} ; # END Namespace
