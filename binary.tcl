namespace eval ::tcltm::binary {
    proc readfile { dir file } {
        set b [open [file normalize [file join $dir [filename $file]]]]
        fconfigure $b -translation binary
        fconfigure $b -encoding binary
        set data [read $b]
        close $b

        return $data
    }

    proc filesize { dir file } {
        return [string length [readfile $dir [filename $file]]]
    }

    proc present { files } {
        foreach f $files {
            if { [file extension [filename $f]] ne ".tcl" } {
                return 1
            }
        }

        return 0
    }

    # Only return the binary files
    proc files { files } {
        set l [list]
        foreach f $files {
            if { [file extension [filename $f]] ne ".tcl" } {
                lappend l $f
            }
        }

        return $l
    }

    proc filename { file } {
        if { [string match "*:*" $file] } {
            set file [lindex [split $file ":"] 0]
        }

        return $file
    }

    proc fileaction { file } {
        set action {}
        if { [string match "*:*" $file] } {
            set action [lindex [split $file ":"] 1]
        }

        return $action
    }
}
