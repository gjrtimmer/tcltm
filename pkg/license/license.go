package license

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gjrtimmer/tcltm/pkg/config"
	"github.com/gjrtimmer/tcltm/pkg/resource"
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
func New(c *config.Config, m *config.Module) (l *License, err error) {
	l.content = make([]byte, 0)

	// Determine license type based on config value
	// First check for a LICENSE file within the InputDirectory
	if _, err = os.Stat(filepath.Join(c.InputDirectory, "LICENSE")); err == nil {
		// LICENSE file present
		// Load LICENSE file
		if l.content, err = ioutil.ReadFile(filepath.Join(c.InputDirectory, "LICENSE")); err != nil {
			return nil, err
		}

		return l, nil
	}

	// Check if license is SPDX
	if _, err = TemplateString(m.License); err == nil {
		// License configuration value is a Template
		//
		// Load Template
		r, _ := resource.Get(fmt.Sprintf("/license/%s", m.License))
		t := template.Must(template.New(m.License).Parse(string(r)))

		aux := struct {
			Authors []config.Author
			Year    int
		}{
			Year: time.Now().Year(),
		}
		if len(m.Authors) > 0 {
			aux.Authors = m.Authors
		}

		// Execute Template
		var b bytes.Buffer
		if err = t.Execute(&b, aux); err != nil {
			return nil, err
		}
		l.content = b.Bytes()
		return l, nil
	}

	// Check if configuration value points to a license file
	// This value is expected to be relative to the input directory
	if _, err := os.Stat(filepath.Join(c.InputDirectory, m.License)); err == nil {
		// Configuration value points to a file.
		// Load license
		if l.content, err = ioutil.ReadFile(filepath.Join(c.InputDirectory, m.License)); err != nil {
			return nil, err
		}
		return l, nil
	}

	// Assume license is inline
	l.content = []byte(m.License)

	return l, nil
}

// EOF
