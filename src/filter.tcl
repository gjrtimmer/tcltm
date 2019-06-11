# filter.tcl
#
#   Handler for filtering source files
#
namespace eval ::tcltm::filter {
    proc line { str key value } {
        regsub -all -- "@${key}@" $str $value str
        return $str
    }

    proc lines { data key value } {
        set lines [list]
        foreach l [split $data "\n"] {
            lappend lines [line $l $key $value]
        }

        return [join $lines "\n"]
    }

    proc multi { data args } {
        set lines [list]
        foreach l [split $data "\n"] {
            set line $l
            foreach {k v} $args {
                set line [line $line $k $v]
            }
            lappend lines $line
        }

        return [join $lines "\n"]
    }

    # filter
    #
    #   List will build a complete filter list from the current
    #   package configuration for the request file.
    #
    # return filter list for specific file
    proc lfile { pkg file } {
        set filter [list]
         # Resolve Global Filter keys
        if { [dict exists $pkg filter] } {
            lappend filter [dict get $pkg filter]
        }

        foreach f [dict get $pkg files] {
            if { [dict exists $pkg filter] && [dict get $f name] eq $file } {
                lappend filter [dict get $f filter]
            }
        }

        return $filter
    }
}
