# config.tcl
#
#   Handler for .tcltm config files
#
namespace eval ::tcltm::config {

    # exists
    #
    #   Check if config exists
    #
    proc exists { dir {cfg .tcltm} } {
        set fname [file normalize [file join $dir $cfg]]
        return [file exists $fname]
    }

    # read
    #
    #   Read config file
    #
    # Return dictionary
    proc load { dir {cfg .tcltm} } {
        set fname [file normalize [file join $dir $cfg]]
        return [::yaml::yaml2dict -file $fname -m:true {1 {true on}} -m:false {0 {false off}}]
    }

    # merge
    #
    #   Merge loaded configuration with commandline options
    #
    # returns merged config
    proc merge { cfg opts } {
        dict set cfg options $opts
        return $cfg
    }

    # parse
    #
    #   Parse the configuration and resolve all environment variables
    #
    # return parsed / resolved configuration
    proc parse { cfg } {
        # Resolve Package version
        set pkgs [list]
        foreach p [dict get $cfg package] {
            # Resolve Global Filter keys
            if { [dict exists $p filter] } {
                set filter [list]
                foreach {k v} [dict get $p filter] {
                    lappend filter "$k [::tcltm::env::resolve $v]"
                }
                dict set p filter $filter
            }

            # Resolve File Filter keys
            set files [list]
            foreach f [dict get $p files] {
                if { [dict exists $f filter] } {
                    set filter [list]
                    foreach {k v} [dict get $f filter] {
                        lappend filter "$k [::tcltm::env::resolve $v]"
                    }
                    dict set f filter $filter
                }
                lappend files $f
            }
            dict set p files $files

            # Resolve Version
            if { [dict exists $p version] } {
                dict set p version [::tcltm::env::resolve [dict get $p version]]
            }

            # Resolve version from pkgIndex.tcl
            if { [dict get $cfg options version-from-index] } {
                set idx [file normalize [file join [dict get $cfg options in] pkgIndex.tcl]]
                if { [file exists $idx] } {
                    set results [::tcltm::scan $idx]
                    foreach {f res} $results {
                        if { $f eq $idx } {
                            foreach pkg $res {
                                if { [dict get $p name] eq [dict get $pkg package] && [dict get $pkg type] eq "ifneeded" } {
                                    dict set p version [dict get $pkg version]
                                }
                            }
                        }
                    }
                }
            }

            # Update Package configuration
            lappend pkgs $p
        }
        dict set cfg package $pkgs
        return $cfg
    }
}
