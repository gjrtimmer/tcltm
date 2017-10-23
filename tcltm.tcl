#!/bin/sh
# the next line restarts using tclsh \
exec tclsh "$0" ${1+"$@"}

if { ![package vsatisfies [package provide Tcl] 8.6] } {puts stdout "Tcl: 8.6 is required"; return}
if { ![package vsatisfies [package require yaml] 0.3.6] } {puts stdout "Yaml: 0.3.6 is required"; return}

#SOURCE#

namespace eval ::tcltm {
    variable config

    proc usage {} {
        puts stdout {tcltm ?options?
    -d DIR, --dir DIR, --directory DIR      Input directory
    -o DIR, --out DIR                       Output directory
    -c                                      Create Tcl target structure
    -s FILE, --scan FILE                    Scan Tcl file for dependencies
    -h, --help                              Show help}
    }

    proc main { args } {
        variable config

        # Parse commandline options
        array set options {directory {} out {} create 0 scan {} help 0 strip 0}
        while { [llength $args] } {
            switch -glob -- [lindex $args 0] {
                -d* -
                --dir* -
                --directory*        {set args [lassign $args - options(directory)]}
                -o* -
                -out*               {set args [lassign $args - options(out)]}
                -c*                 {set options(create) 1; set args [lrange $args 1 end]}
                -s* -
                --scan*             {set args [lassign $args - options(scan)]}
                --strip-comments*   {set options(strip) 1; set args [lrange $args 1 end]}
                -h* -
                --help*             {set options(help) 1; set args [lrange $args 1 end]}
                default             {break}
            }
        }

        if { $options(help) } {
            usage
            exit 1
        }

        if { [string length $options(scan)] > 0 } {
            #puts stdout [file normalize $options(scan)]
            #if { ![file exists [file normalize $options(scan)]] } {
            #    puts stdout "File [file normalize $options(scan)] does not exists";flush stdout
            #    exit 1
            #}
            set f [file normalize [lindex $options(scan) 0]]
            if { ![file exists $f] } {
                puts stdout "File '$f' does not exists"
                exit 1
            }
            
            puts stdout "Scanning: $f"
            set b [open $f]
            fconfigure $b -translation binary
            fconfigure $b -encoding binary
            set data [read $b]
            close $b

            foreach line [split $data "\n"] {
                if { [regexp {package require(?:[[:blank:]]+)([_[:alpha:]][:_[:alnum:]]*)((?:[[:blank:]]+)?(?:(\d+\.)?(\d+\.)?(\*|\d+))?)} $line match pkg ver] } {
                    puts stdout "$pkg [string trim $ver]"
                }
            }

            exit
        }

        if { [string length $options(directory)] == 0 } {
            puts stdout "No input directory provided"
            puts stdout "  => Using current working directory \[[file normalize [pwd]]\]"
            set options(directory) [file normalize [pwd]]
            flush stdout
        }

        if { [string length $options(out)] == 0 } {
            puts stdout "No output directory provided"
            puts stdout "  => Using current working directory \[[file normalize [pwd]]\]"
            set options(out) [file normalize [pwd]]
            flush stdout
        }

        # Check for configuration
        if { ![::tcltm::config::exists $options(directory)] } {
            puts "Missing .tcltm specification"; exit 1
        }
        
        set config [::tcltm::config::load $options(directory)]

        foreach pkg [dict get $config Package] {
            set pkgcontent [list]

            # Check for license
            if { ![dict exists $pkg License] || [string length [dict get $pkg License]] == 0 } {
                if { [::tcltm::license::exists $options(directory)] } {
                    dict set pkg License [::tcltm::license::load $options(directory)]
                }
            }

            # License
            if { [string length [dict get $pkg License]] > 0 } {
                lappend pkgcontent {*}[::tcltm::license::format [dict get $pkg License]]
                lappend pkgcontent [::tcltm::markup::nl]
            }

            # Header
            lappend pkgcontent [::tcltm::markup::comment "TCLTM HEADER BEGIN"]
            lappend pkgcontent [::tcltm::markup::meta "PACKAGE" [dict get $pkg Name] [dict get $pkg Version]]
            foreach key {Summary Description Tcl} {
                if { [dict exists $pkg $key] && [string length [dict get $pkg $key]] > 0 } {
                    if { $key eq "Description" } {
                        foreach line [split [dict get $pkg $key] "\n"] {
                            lappend pkgcontent [::tcltm::markup::meta "DESCRIPTION" $line]
                        }
                    } else {
                        lappend pkgcontent [::tcltm::markup::meta $key [dict get $pkg $key]]
                    }
                }
            }
            
            # Header => Dependencies
            if { [dict exists $pkg Dependencies] && [string length [dict get $pkg Dependencies]] > 0 } {
                foreach r [dict get $pkg Dependencies] {
                    lappend pkgcontent [::tcltm::markup::meta "REQUIRE" $r]
                }
            }
            lappend pkgcontent [::tcltm::markup::comment "TCLTM HEADER END"]
            lappend pkgcontent [::tcltm::markup::nl]

            # Dependencies
            lappend pkgcontent [::tcltm::markup::script {
if { ![package vsatisfies [package provide Tcl] %s] } {
    return -code error "Unable to load module '%s' Tcl: '%s' is required"
}
            } [dict get $pkg Tcl] [dict get $pkg Name] [dict get $pkg Tcl]]

            if { [dict exists $pkg Dependencies] && [string length [dict get $pkg Dependencies]] > 0 } {
                foreach r [dict get $pkg Dependencies] {
                    lappend pkgcontent [::tcltm::markup::script {package require %s} $r]
                }
                lappend pkgcontent [::tcltm::markup::nl]
            }

            # Package declaration
            lappend pkgcontent [::tcltm::markup::script {package provide %s %s} [dict get $pkg Name] [dict get $pkg Version]]
            lappend pkgcontent [::tcltm::markup::nl]

            # Detect if binary loader needs to be activated
            if { [::tcltm::binary::present [dict get $pkg Files]] } {
                lappend pkgcontent [::tcltm::markup::comment "TCLTM BINARY LOADER START"]
                lappend pkgcontent [::tcltm::markup::script {
catch {set tmpdir $::env(TMP)}
catch {set tmpdir $::env(TMPDIR)}
catch {set tmpdir $::env(TEMP)}
set bin [open [info script] {RDONLY BINARY}]
set bindata [read $bin]
close $bin
unset -nocomplain bin

set binaryIndex [string first \\u001A $bindata]
incr binaryIndex}]
                
                foreach f [::tcltm::binary::files [dict get $pkg Files]] {
                    lappend pkgcontent [::tcltm::markup::script {
set tmpBinary [file normalize [file join $tmpdir %s]]
set fh [open $tmpBinary w]
fconfigure $fh -translation binary
fconfigure $fh -encoding binary
puts -nonewline $fh [string range $bindata $binaryIndex [incr binaryIndex %s]-1]
flush $fh
close $fh} [::tcltm::binary::filename $f] [::tcltm::binary::filesize $options(directory) [::tcltm::binary::filename $f]]]
                    
                    # Add auto-load code
                    if { [::tcltm::binary::fileaction $f] eq "load" } {
                        lappend pkgcontent [::tcltm::markup::script {
load $tmpBinary
catch {file delete -force $tmpBinary}
unset -nocomplain fh}]
                    } else {
                        # No auto-loading of libraries
                        # assume the user wants to handle the binary themselfs.
                        # add extract binaries to variable
                        # so the user can use it within the init script
                        lappend pkgcontent [::tcltm::markup::script {lappend binFiles $tmpBinary}]
                    }
                    lappend pkgcontent [::tcltm::markup::script {unset -nocomplain tmpBinary}]
                }

                # Auto auto unload variables
                lappend pkgcontent [::tcltm::markup::script {unset -nocomplain bindata binaryIndex}]

                lappend pkgcontent [::tcltm::markup::comment "TCLTM BINARY LOADER END"]
                lappend pkgcontent [::tcltm::markup::nl]
            }

            # Load InitScript into content
            # InitScript should be placed after the binary loader
            # This allows a InitScript to handle the binary content
            # for example a InitScript which auto-extracts binary
            # attached content
            #
            # the binary content can be loaded from the index
            # within variable 'binaryIndex'
            if { [dict exists $pkg InitScript] && [string length [dict get $pkg InitScript]] > 0} {
                lappend pkgcontent [::tcltm::markup::comment "TCLTM INIT BEGIN"]
                if { [llength [split [dict get $pkg InitScript] "\n"]] == 1 } {
                    if { [string match "*.tcl" [lindex [split [dict get $pkg InitScript] "\n"] 0]] } {
                        set initfile [lindex [split [dict get $pkg InitScript] "\n"] 0]
                        foreach line [split [::tcltm::binary::readfile $options(directory) [::tcltm::binary::filename $initfile]] "\n"] {
                            if { $options(strip) && [::tcltm::markup::iscomment $line] } {
                                # Ignore line
                            } else {
                                lappend pkgcontent $line
                            }
                        }
                    }
                } else {
                    # assume embedded Tcl code in InitScript specification
                    foreach line [split [dict get $pkg InitScript] "\n"] {
                        if { $options(strip) && [::tcltm::markup::iscomment $line] } {
                            # Ignore line
                        } else {
                            lappend pkgcontent [::tcltm::markup::script $line]
                        }
                    }
                }
                lappend pkgcontent [::tcltm::markup::comment "TCLTM INIT END"]
                lappend pkgcontent [::tcltm::markup::nl]
            }

            # Add script files
            lappend pkgcontent [::tcltm::markup::comment "TCLTM SCRIPT SECTION BEGIN"]
            foreach f [dict get $pkg Files] {
                if { [file extension [::tcltm::binary::filename $f]] eq ".tcl" } {
                    foreach line [split [::tcltm::binary::readfile $options(directory) [::tcltm::binary::filename $f]] "\n"] {
                        if { $options(strip) && [::tcltm::markup::iscomment $line] } {
                            # Ignore line
                        } else {
                            if { ![regexp {^(?:([[:blank:]]+)?)package provide*} $line] && ![regexp {^(?:([[:blank:]]+)?)package require.*$} $line] } {
                                lappend pkgcontent $line
                            } 
                        }
                    }
                }
            }
            lappend pkgcontent [::tcltm::markup::comment "TCLTM SCRIPT SECTION END"]

            if { [::tcltm::binary::present [dict get $pkg Files]] } {
                lappend pkgcontent [::tcltm::markup::nl]
                lappend pkgcontent [::tcltm::markup::comment "TCLTM BINARY SECTION BEGIN"]
            }

            # Write Tcl Module            
            set filename [format {%s-%s.tm} [dict get $pkg Name] [dict get $pkg Version]]
            set filepath [file normalize [file join [file dirname [file normalize $options(out)]] $filename]]
            puts stdout "Package: $filename"
            set fh [open $filepath w]
            fconfigure $fh -translation lf
            puts $fh [join $pkgcontent "\n"]
            puts -nonewline $fh "\u001A"
            fconfigure $fh -translation binary
            foreach file [::tcltm::binary::files [dict get $pkg Files]] {
                puts $fh [::tcltm::binary::readfile $options(directory) $file]
            }
            close $fh
        }
    }
}

::tcltm::main {*}$::argv

# EOF