namespace eval ::tcltm::license {
    variable filename LICENSE

    proc exists { dir } {
        variable filename

        set fname [file normalize [file join $dir $filename]]
        return [file exists $fname]
    }

    proc load { dir } {
        variable filename

        set fname [file normalize [file join $dir $filename]]
        set fh [open $fname RDONLY]
        set data [read $fh]
        close $fh

        return $data
    }

    proc format { data } {
        set license [list]
        lappend license $::tcltm::markup::divider "#"
        foreach line [split $data "\n"] {
            lappend license [::tcltm::markup::comment $line]
        }
        lappend license "#" $::tcltm::markup::divider

        return $license
    }
}
