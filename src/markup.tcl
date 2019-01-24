# markup.tcl
#
#   Markup handler for tcltm
#
namespace eval ::tcltm::markup {
    variable divider [string repeat "#" 80]

    proc comment { n args } {
        set line {}

        if { [llength $args] } {
            set line [format {# %s %s} $n [join $args]]
        } else {
            set line [format {# %s} $n]
        }

        return $line
    }

    proc iscomment { line } {
        if { [string index $line 0] eq "#" } {
            return 1
        }

        return 0
    }

    proc nl {} {
        return {}
    }

    proc meta { n args } {
        if { [llength $args] } {
            set line [format {# %s: %s} [string toupper $n] [join $args]]
        } else {
            set line [format {# %s} [string toupper $n]]
        }

        return $line
    }

    proc script { body args } {
        regsub -all "\n$" $body {} body
        return [string trimleft [format "[subst -nocommands -novariables $body]" {*}$args] "\n"]
        # set script [string trimleft [format "[subst -nocommands -novariables $body]" {*}$args] "\n"]

        # set lines [list]
        # foreach l [split $script "\n"] {
        #     lappend lines [string trimleft $l]
        # }

        # return [join $lines "\n"]
    }
}
