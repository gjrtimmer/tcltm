package license

// Type defines the different type of used license(s)
type Type uint8

const (
	// TypeTemplate is a license which is one of the embedded
	// template licenses within tcltm
	TypeTemplate Type = iota

	// TypeFile defines a license which is defined within a file
	TypeFile

	// TypeInline is a license which is provided inline within the configuration
	TypeInline
)

// EOF
