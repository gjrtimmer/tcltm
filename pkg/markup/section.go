//go:generate enumer -type=Section -output=section_type.go -linecomment

package markup

import (
	"bytes"
	"fmt"
)

const (
	// SectionBegin defines the begin of a section
	SectionBegin = "TCLTM %s BEGIN"

	// SectionEnd defines the end of section
	SectionEnd = "TCLTM %s END"
)

// Section defines the different sections within
// a generated module.
//
// Each section has a begin and an end.
type Section uint8

const (
	// SectionHeader defines the header section
	SectionHeader Section = iota // HEADER

	// SectionBinaryLoader defines the section which holds
	// the binary loader
	SectionBinaryLoader // BINARY LOADER

	// SectionScript defines the script section
	SectionScript // SCRIPT

	// SectionBootstrapScript defines the bootstrap section
	SectionBootstrapScript // BOOTSTRAP

	// SectionInitScript defines the initscript section
	SectionInitScript // INIT

	// SectionFinalizeScript defines the finalize script section
	SectionFinalizeScript // FINALIZE
)

// WriteSection will write a section line to the buffer
// where s is the Section to write and t is either
// SectionBegin or SectionEnd
func WriteSection(b *bytes.Buffer, s Section, t string) (int, error) {
	return Comment(b, fmt.Sprintf(t, s.String()))
}

// EOF
