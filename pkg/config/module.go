//go:generate enumer -type=FileType -output=type_file.go -yaml -trimprefix=FileType -linecomment
//go:generate enumer -type=FileAction -output=type_action.go -yaml -trimprefix=FileAction -linecomment

package config

// Module defines the build configuration for a single module.
type Module struct {
	Include      *Include          `yaml:"include,omitempty"`
	Exclude      *Exclude          `yaml:"exclude,omitempty"`
	Output       *Output           `yaml:"output,omitempty"`
	Name         string            `yaml:"name"`
	Version      string            `yaml:"version"`
	Authors      []Author          `yaml:"authors,omitempty"`
	Tcl          string            `yaml:"tcl"`
	Interpreter  string            `yaml:"interp,omitempty"`
	Summary      string            `yaml:"summary,omitempty"`
	Description  string            `yaml:"description,omitempty"`
	License      string            `yaml:"license,omitempty"`
	Dependencies []string          `yaml:"dependencies,omitempty"`
	Meta         map[string]string `yaml:"meta,omitempty"`
	Filter       map[string]string `yaml:"filter,omitempty"`
	Extension    string            `yaml:"extension,omitempty"`
	FinalName    string            `yaml:"finalname,omitempty"`
	Bootstrap    string            `yaml:"bootstrap,omitempty"`
	InitScript   string            `yaml:"init,omitempty"`
	Files        []File            `yaml:"files"`
}

// Author definition
type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email,omitempty"`
}

// File definition
type File struct {
	Name      string            `yaml:"name"`
	Type      FileType          `yaml:"type"`
	Action    FileAction        `yaml:"action"`
	Filter    map[string]string `yaml:"filter,omitempty"`
	Filtering bool              `yaml:"filtering,omitempty"`
	Content   []byte            `yaml:"-"`
}

// FileType defines the possible file types
type FileType uint8

const (
	// FileTypeScript represents a script
	FileTypeScript FileType = iota // script

	// FileTypeBinary represents a binary file
	FileTypeBinary // binary
)

// FileAction defines the possible actions for an embedded file
type FileAction uint8

const (
	// FileActionNone defines no file action
	FileActionNone FileAction = iota // none

	// FileActionRun defines to run the specific file
	FileActionRun //run

	// FileActionLoad defines to load the specific file
	// can be used to 'load' a library (.dll/.so)
	FileActionLoad // load
)

// EOF
