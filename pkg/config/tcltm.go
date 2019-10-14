package config

// TclTM defines the global configuration of the module builder
type TclTM struct {
	// Include configuration
	Include *Include `yaml:"include,omitempty"`

	// Exclude configuration
	Exclude *Exclude `yaml:"exclude,omitempty"`

	// Output configuration
	Output *Output `yaml:"output,omitempty"`
}

// Output defines the output configuration
type Output struct {
	// Repository will when enabled write the modules defined within the configuration
	// to an output directory conforming to the default Tcl TM Repository format.
	//
	// default: false
	Repository *bool `yaml:"repository,omitempty"`

	// InteractiveLoader will reconfigure the binary loader to
	// only be executed when the module is loaded within an
	// tcl_interactive state when enabled.
	//
	// default: false
	InteractiveLoader *bool `yaml:"interactive-loader,omitempty"`
}

// Exclude defines the configuration what to exclude from
// the TclTM output
type Exclude struct {
	// Comments will configure if comments should be stripped
	// from the ouput modules.
	//
	// default: false
	Comments *bool `yaml:"comments,omitempty"`

	// ResourcePrefix will strip the provided prefix of any included
	// resources within the module.
	//
	// When the .tcltm configuration uses relative or absolutes paths
	// for its resources its path gets included in the output.
	//
	// default: ""
	ResourcePrefix *string `yaml:"resource-prefix,omitempty"`

	// ResourcePath will strip the complete prefix path of any included
	// resources. When enabled will strip the provided resource directory
	// prefix from the included resources.
	//
	//  Example Configuration:
	ResourcePath *bool `yaml:"resource-path,omitempty"`

	// Dependencies will control is the provided dependencies of a module
	// are to be included with a `package require` command.
	//
	// If this property is false, the no `package require` commands are written
	// to the generated module
	Dependencies *bool `yaml:"dependencies,omitempty"`

	// Provide will control if the `package provide` command is written to the
	// generated module.
	//
	// The primary function of tcltm is to build Tcl modules (.tm)
	// These modules require a `package provide` command.
	//
	// Because tcltm can also be used to generate custom modules or
	// packages with embedded resources the default written `package provide`
	// is not always needed.
	Provide *bool `yaml:"provide,omitempty"`

	// SatisfyTcl will control if the default satisfytcl command is written to
	// the genereated output.
	//
	// By default tcltm will write the following command to ensure
	// the Tcl version is satisfied.
	//
	//  if { ![package vsatisfies [package provide Tcl] %s] } {
	//      return -code error "Unable to load module '%s' Tcl: '%s' is required"
	//  }
	SatisfyTcl *bool `yaml:"satisfy-tcl,omitempty"`
}

// Include defines the configuration what to include or preserve
// from the included resources files
type Include struct {
	// Require will ensure that the 'package require' commands will
	// be preserved from the included resources.
	Require *bool `yaml:"require,omitempty"`
}

// EOF
