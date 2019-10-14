package license

import (
	"os"
	"path/filepath"

	"github.com/gjrtimmer/tcltm/pkg/config"
)

// License represents a module license
type License struct {
	content []byte
}

// New create a new license object
//
// Config is required for InputDirectory to find the license
// if this is a file.
//
// A new license is constructed based upon the provided module
func New(c *config.Config, m *config.Module) *License {
	l := &License{}

	// Determine license type based on config value
	// First check for a LICENSE file within the InputDirectory
	if _, err := os.Stat(filepath.Join(c.InputDirectory, "LICENSE")); err == nil {
		// LICENSE file present
		// Load LICENSE file
	}

	// 1) Check if license is SPDX
	// 2) Check if license is file
	// 3) !1 && !2 => License is inline
	// if spdx, err := SPDXString(m.License); err == nil {
	// 	// License value is SPDX
	// 	// Load embedded license
	// }

	//

	return l
}

// EOF
