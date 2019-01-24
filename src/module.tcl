namespace eval ::tcltm::module {
    variable config     [dict create]
    variable content    [list]

    # Initialize New Module
    # Used when building multiple package from single configuration
    proc new { cfg } {
        variable config $cfg
        variable content

        set config $cfg
        set content [list]

        lappend content [::tcltm::markup::comment "Tcl Module Generated by tcltm; DO NOT EDIT"]
        lappend content [::tcltm::markup::nl]

        return -code ok
    }

    proc pkgcfg { pkg } {
        variable config
        foreach p [dict get $config package] {
            if { [dict get $p name] eq $pkg } {
                return $p
            }
        }

        return -code ok
    }

    proc write { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        # Set extension
        set ext .tm
        if { [dict exists $cfg extension] } {
            set ext [dict get $cfg extension]
        }

        # Build module filename
        set filename [format {%s-%s%s} [dict get $cfg name] [dict get $cfg version] $ext]
        if { [dict exists $cfg finalname] && [string length [dict get $cfg finalname]] > 0 } {
            set filename [dict get $cfg finalname]
        }

        # Tcl Module '::' Fix
        # For a Tcl module '::' within the package name
        # is replaced by the os seperator
        regsub -all -- {::} $filename {/} filename

        # Output Directory
        set filepath [file normalize [file join [file normalize [dict get $config options out]] $filename]]
        if { [dict get $config options repo] } {
            set tcldir "tcl[lindex [split [dict get $cfg tcl] "."] 0]"
            set outdir [file normalize [file join [dict get $config options out] $tcldir [dict get $cfg tcl]]]
            # Convert to Tcl: 8.5
            try {
                file mkdir $outdir
            } on error err {
                puts stdout "Failed to create output directory \[$outdir\]"
            }
            set filepath [file join $outdir $filename]
        }

        # Create dirs for '::' replacement
        # TODO: Convert to Tcl 8.5
        file mkdir [file dirname $filepath]

        # Place binary comment marker
        if { [::tcltm::binary::present [dict get $cfg files]] } {
            lappend content [::tcltm::markup::nl]
            lappend content [::tcltm::markup::comment "BINARY SECTION"]
        }

        # Write Module
        set fh [open $filepath w]
        fconfigure $fh -translation lf
        set lines [join $content "\n"]

        # Strip Double Empty lines
        regsub -all -- {\n\n\n+} $lines "\n\n" lines
        puts $fh $lines

        # Write Binary Marker
        if { [::tcltm::binary::present [dict get $cfg files]] } {
            puts -nonewline $fh "\u001A"
            fconfigure $fh -translation binary

            # Encode binary files
            set binfiles [::tcltm::binary::files [dict get $cfg files]]
            foreach f $binfiles {
                puts stdout "Encoding: [dict get $f name]"
                puts $fh [::tcltm::binary::readfile [dict get $config options in] [dict get $f name]]
            }
        }

        # Finalize Module
        close $fh
        puts stdout "Module: $filename \[$filepath\]"

        return -code ok
    }

    proc license { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        # If no license is configured
        # try to load license from default LICENSE file
        if { ![dict exists $cfg license] || [string length [dict get $cfg license]] == 0 } {
            if { [::tcltm::license::exists [dict get $config options in]] } {
                dict set cfg license [::tcltm::license::load [dict get $config options in]]
            }
        }

        if { [string length [dict get $cfg license]] > 0 } {
            # License configured
            if { [llength [split [dict get $cfg license] "\n"]] == 1 } {
                # No multiline license configured in configuration
                # try to load license file
                dict set cfg license [::tcltm::license::load [dict get $config options in] [dict get $cfg license]]
            }
        }

        # Append license to module
        if { [dict exists $cfg license] && [string length [dict get $cfg license]] > 0 } {
            lappend content {*}[::tcltm::license::format [dict get $cfg license]]
            lappend content [::tcltm::markup::nl]
        }

        return -code ok
    }

    proc header { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        # Header Start
        lappend content [::tcltm::markup::comment "TCLTM HEADER BEGIN"]

        # Name
        # Version
        # Summary
        # Description
        # Tcl+
        foreach key {name version summary description Tcl} {
            if { [dict exists $cfg $key] && [string length [dict get $cfg $key]] > 0 } {
                if { [string tolower $key] eq "description" } {
                    foreach line [split [dict get $cfg $key] "\n"] {
                        lappend content [::tcltm::markup::meta "DESCRIPTION" $line]
                    }
                } else {
                    lappend content [::tcltm::markup::meta $key [dict get $cfg $key]]
                }
            }
        }

        # Dependencies
        if { [dict exists $cfg dependencies] && [string length [dict get $cfg dependencies]] > 0 } {
            foreach r [dict get $cfg dependencies] {
                lappend content [::tcltm::markup::meta "REQUIRE" $r]
            }
        }

        # Add embedded binary resources to HEADER
        set files [list]
        set bidx 0
        for {set fidx 0} {$fidx < [llength [dict get $cfg files]]} {incr fidx} {
            set fcfg [lindex [dict get $cfg files] $fidx]

            # Only process binary files
            if { ![dict exists $fcfg type] } {
                dict set fcfg type "script"
            }
            if { [string toupper [dict get $fcfg type]] eq "BINARY" } {
                # File Index
                dict set fcfg id $bidx
                incr bidx

                # Encode file to header
                set enc [::tcltm::binary::encode [dict get $config options in] [dict get $fcfg name]]
                set fcfg [list {*}$fcfg {*}$enc]

                # Create Header
                set header [format {ID %s NAME %s SIZE %s HASH %s} \
                    [dict get $fcfg id] \
                    [dict get $fcfg name] \
                    [dict get $fcfg size] \
                    [dict get $fcfg hash] \
                ]

                # Append action if provided
                if { [dict exists $fcfg action] } {
                    append header " ACTION [dict get $fcfg action]"
                }

                # Append target if exists
                if { [dict exists $fcfg target] } {
                    append header " TARGET [dict get $fcfg target]"
                }

                lappend content [::tcltm::markup::meta "RESOURCE" [format "{%s}" $header]]
            }

            lappend files $fcfg
        }

        # Header End
        lappend content [::tcltm::markup::comment "TCLTM HEADER END"]

        return -code ok
    }

    proc satisfy-tcl-version { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        if { ![dict get $config options exclude-satisfy-tcl] } {
            lappend content [::tcltm::markup::nl]
            lappend content [::tcltm::markup::script {
if { ![package vsatisfies [package provide Tcl] %s] } {
    return -code error "Unable to load module '%s' Tcl: '%s' is required"
}
} [dict get $cfg tcl] [dict get $cfg name] [dict get $cfg tcl]]
        }

        return -code ok
    }

    proc deps { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        if { ![dict get $config options exclude-deps] } {
            if { [dict exists $cfg dependencies] && [string length [dict get $cfg dependencies]] > 0 } {
                lappend content [::tcltm::markup::nl]
                foreach r [dict get $cfg dependencies] {
                    lappend content [::tcltm::markup::script {package require %s} $r]
                }
            }
        }

        return -code ok
    }

    proc script { pkg type } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        if { [dict exists $cfg $type] && [string length [dict get $cfg $type]] > 0 } {
            lappend content [::tcltm::markup::nl]
            lappend content [::tcltm::markup::comment "TCLTM [string toupper $type] BEGIN"]
            if { [llength [split [dict get $cfg $type] "\n"]] == 1 } {
                if { [string match "*.tcl" [lindex [split [dict get $cfg $type] "\n"] 0]] } {
                    set bfile [lindex [split [dict get $cfg $type] "\n"] 0]
                    foreach line [split [::tcltm::binary::readfile [dict get $config options in] [::tcltm::binary::filename $bfile]] "\n"] {
                        if { [dict get $config options strip] && [::tcltm::markup::iscomment $line] } {
                            # Ignore line
                        } else {
                            lappend content $line
                        }
                    }
                } else {
                    lappend content [::tcltm::markup::script [dict get $cfg $type]]
                }
            } else {
                # assume embedded Tcl code in InitScript specification
                 foreach line [split [dict get $cfg $type] "\n"] {
                    if { [dict get $config options strip] && [::tcltm::markup::iscomment $line] } {
                        # Ignore line
                    } else {
                        lappend content [::tcltm::markup::script $line]
                    }
                }
            }
            lappend content [::tcltm::markup::comment "TCLTM [string toupper $type] END"]
        }

        return -code ok
    }

    proc code { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        lappend content [::tcltm::markup::nl]
        lappend content [::tcltm::markup::comment "TCLTM SCRIPT SECTION BEGIN"]
        foreach f [dict get $cfg files] {
            set inc 0
            if { [file extension [::tcltm::binary::filename [dict get $config options in] [dict get $f name]]] eq ".tcl" } {
                set inc 1
            } elseif { [dict exists $f type] && [string tolower [dict get $f type]] eq "script" } {
                set inc 1
            }

            set filter [list]
            if { [dict exists $f filtering] && [dict get $f filtering] } {
                set filter {*}[::tcltm::filter::lfile $cfg [dict get $f name]]
            }

            if { $inc } {
                set ignore(block) 0
                set ignore(next) 0
                foreach line [split [::tcltm::binary::readfile [dict get $config options in] [dict get $f name]] "\n"] {
                    if { [string match {*TCLTM*IGNORE*BEGIN*} [string toupper $line]] } {
                        set ignore(block) 1
                        continue
                    }
                    if { [string match {*TCLTM*IGNORE*END*} [string toupper $line]] } {
                        set ignore(block) 0
                        continue
                    }
                    if { $ignore(block) } {
                        continue
                    }
                    if { [string match {*TCLTM*IGNORE*NEXT*} [string toupper $line]] } {
                        set ignore(next) 1
                        continue
                    }
                    if { $ignore(next) } {
                        set ignore(next) 0
                        continue
                    }
                    if { [string match {*TCLTM*IGNORE*} [string toupper $line]] } {
                        continue
                    }
                    if { [dict get $config options strip] && [::tcltm::markup::iscomment $line] } {
                        # Ignore line
                    } else {
                        if { ![regexp {^(?:([[:blank:]]+)?)package provide*} $line] && ![regexp {^(?:([[:blank:]]+)?)package require.*$} $line] } {

                            # Filter Lines
                            if { [dict exists $f filtering] && [dict get $f filtering] } {
                                foreach {k v} {*}$filter {
                                    set line [::tcltm::filter::line $line $k $v]
                                }
                            }

                            lappend content $line
                        }
                    }
                }
            }
        }
        lappend content [::tcltm::markup::comment "TCLTM SCRIPT SECTION END"]

        return -code ok
    }

    proc pkg-provide { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        if { ![dict get $config options exclude-provide] } {
            lappend content [::tcltm::markup::nl]
            lappend content [::tcltm::markup::script {package provide %s %s} [dict get $cfg name] [dict get $cfg version]]
        }

        return -code ok
    }

    proc binaryloader { pkg } {
        variable config
        variable content
        variable cfg [pkgcfg $pkg]

        if { [::tcltm::binary::present [dict get $cfg files]] } {
            lappend content [::tcltm::markup::nl]
            lappend content [::tcltm::markup::comment "TCLTM BINARY LOADER BEGIN"]
            lappend content [::tcltm::markup::script $::tcltm::loader::script]
            lappend content [::tcltm::markup::comment "TCLTM BINARY LOADER END"]
        }

        return -code ok
    }
}
