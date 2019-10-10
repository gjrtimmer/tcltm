# license.tcl
#
#   Handler for project LICENSE
#
namespace eval ::tcltm::license {
    proc exists { dir {filename LICENSE} } {
        set fname [file normalize [file join $dir $filename]]
        return [file exists $fname]
    }

    proc load { dir {filename LICENSE} } {
        set fname [file normalize [file join $dir $filename]]
        set fh [open $fname RDONLY]
        set data [read $fh]
        close $fh

        return $data
    }

    # Format proc
    proc format { data } {
        set license [list]
        lappend license $::tcltm::markup::divider
        foreach line [split $data "\n"] {
            if { $line eq {} } {
                lappend license "#"
            } else {
                lappend license [::tcltm::markup::comment $line]
            }
        }
        lappend license $::tcltm::markup::divider

        return $license
    }
}
