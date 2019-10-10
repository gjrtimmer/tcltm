namespace eval ::tcltm::binary {
    proc readfile { dir file } {
        set b [open [file normalize [filename $dir $file]]]
        fconfigure $b -translation binary
        fconfigure $b -encoding binary
        set data [read $b]
        close $b

        return $data
    }

    proc filesize { dir file } {
        return [string length [readfile $dir $file]]
    }

    proc filename { dir file } {
        set f $file
        if { [string match {*\**} $f] } {
            set f [glob -directory $dir $f]
            return $f
        }
        return [file normalize [file join $dir $file]]
    }

    proc hash { dir file } {
        return [::sha1::sha1 -hex -file [filename $dir $file]]
    }

    proc encode { dir file } {
        set info [dict create]
        dict set info size [filesize $dir $file]
        dict set info hash [hash $dir $file]

        return $info
    }

    proc present { flist } {
        return [expr {[llength [files $flist]] > 0 ? 1 : 0}]
    }

    proc files { flist } {
        set filelist [list]

        for {set fidx 0} {$fidx < [llength $flist]} {incr fidx} {
            set fcfg [lindex $flist $fidx]

            # Only process binary files
            if { ![dict exists $fcfg type] } {
                dict set fcfg type "script"
            }
            if { [string toupper [dict get $fcfg type]] eq "BINARY" } {
                lappend filelist $fcfg
            }
        }

        return $filelist
    }
}
