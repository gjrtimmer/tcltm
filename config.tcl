# config.tcl
#
#   Handler for .tcltm config files
#
namespace eval ::tcltm::config {

    # exists
    #
    #   Check if config exists
    #
    proc exists { dir } {
        set fname [file normalize [file join $dir .tcltm]]
        return [file exists $fname]
    }

    # read
    #
    #   Read config file
    #
    # Return dictionary
    proc load { dir } {
        set fname [file normalize [file join $dir .tcltm]]
        set fh [open $fname RDONLY]
        set data [read $fh]
        close $fh

        return [::yaml::yaml2dict $data]
    }
}
