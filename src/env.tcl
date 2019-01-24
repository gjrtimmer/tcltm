namespace eval ::tcltm::env {
    proc resolve { val } {
        set v {}

        # Check if value needs to be resolved from environment
        if { [string tolower [string range $val 0 3]] eq "env:" } {
            # Resolvement required
            set l [split $val ":"]

            if { [info exists ::env([lindex $l 1])] } {
                set v $::env([lindex $l 1])
            } elseif { [llength $l] == 3 } {
                set v [lindex $l 2]
            } else {
                error "environment variable '[lindex $l 1]' does not exists"
            }
        }

        return $v
    }
}
