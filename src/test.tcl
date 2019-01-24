package require tcltest
namespace import ::tcltest::*

# Configure Tests
configure -testdir src

# Hook to determine if any of the tests failed. Then we can exit with
# proper exit code: 0=all passed, 1=one or more failed
proc tcltest::cleanupTestsHook {} {
    variable numTests
    set ::exitCode [expr {$numTests(Failed) > 0}]
}

# Run Tests
runAllTests
exit $exitCode
