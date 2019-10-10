namespace eval ::tcltm {
    proc scan { args } {
        set results [dict create]
        set f [file normalize [lindex $args 0]]
        if { ![file exists $f] } {
            puts stdout "File '$f' does not exists"
            exit 1
        }

        set files $f
        if { [file isdirectory $f] } {
            set files [glob -nocomplain -directory $f -types f -- *.tcl]
        }

        foreach f $files {
            set res [dict create]

            set b [open $f]
            fconfigure $b -translation binary
            fconfigure $b -encoding binary
            set data [read $b]
            close $b

            set pkgs [list]
            foreach line [split $data "\n"] {
                set r [dict create]
                if { [regexp {package (provide|require|ifneeded)(?:[[:blank:]]+)([_[:alpha:]][:_[:alnum:]]*)(?:\])?((?:[[:blank:]]+)?(?:(\d+\.)?(\d+\.)?(\*|\d+))?)} $line -> type pkg ver] } {
                    dict set r type $type
                    dict set r package $pkg
                    dict set r version [string trim $ver]
                    lappend pkgs $r
                }
            }

            dict set results $f $pkgs
        }

        return $results
    }
}
