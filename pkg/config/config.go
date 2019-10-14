package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gitlab.timmertech.nl/go/interpolate"
	"gopkg.in/yaml.v2"
)

var (
	// ErrResolveEnvFailed defines the failure to resolve
	// a variable from the environment
	ErrResolveEnvFailed = errors.New("failed to resolve variable from environment")

	// ErrPropertyRequired is returned when a required property is missing
	ErrPropertyRequired = errors.New("required property not provided")

	// ErrModuleNameRequired is returned when a module has no property 'name'
	ErrModuleNameRequired = errors.New("module name is required")

	// ErrModuleVersionRequired is returned when a module has no property 'version'
	ErrModuleVersionRequired = errors.New("module version is required")

	// ErrModuleTclVersionRequired is returned when a module has no property 'tcl'
	ErrModuleTclVersionRequired = errors.New("module tcl version is required")
)

var (
	// Environment holds the variables from the environment
	// This is required for resolving the variable in the configuration
	//
	// Environment is exported to facilitate an override
	//
	// Default loads the OS variables
	Environment interpolate.Env
)

func init() {
	Environment = interpolate.NewSliceEnv(os.Environ())
}

// Config represents a tcltm Configuration
type Config struct {
	// Configuration Version Number
	Version string

	// InputDirectory
	InputDirectory string

	// OutputDirectory
	OutputDirectory string

	// TclTM Configuration
	TclTM TclTM

	// Module Configuration
	Modules []Module
}

// Load loads a package configuration and returns
// a configuration object
func Load(filepath string, strict bool) (*Config, error) {
	cfg := &Config{}

	// Check if config exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, err
	}

	// Read Configuration file
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Parse Config
	if err = yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}

	// Resolve variable in configuration
	if err = cfg.Resolve(strict); err != nil {
		return nil, err
	}

	if err = cfg.Validate(strict); err != nil {
		return nil, err
	}

	return cfg, nil
}

// UnmarshalYAML custom decoder tcltm configuration
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Version string   `yaml:"version"`
		TclTM   TclTM    `yaml:"tcltm,omitempty"`
		Modules []Module `yaml:"modules"`
	}

	if err := unmarshal(&aux); err != nil {
		return err
	}

	// Overrides
	for m := range aux.Modules {
		// Include
		if aux.Modules[m].Include == nil {
			aux.Modules[m].Include = aux.TclTM.Include
		}

		if aux.Modules[m].Include.Require == nil {
			*aux.Modules[m].Include.Require = *aux.TclTM.Include.Require
		}

		// Exclude
		if aux.Modules[m].Exclude == nil {
			aux.Modules[m].Exclude = aux.TclTM.Exclude
		}

		if aux.Modules[m].Exclude.Comments == nil {
			*aux.Modules[m].Exclude.Comments = *aux.TclTM.Exclude.Comments
		}

		if aux.Modules[m].Exclude.ResourcePrefix == nil {
			*aux.Modules[m].Exclude.ResourcePrefix = *aux.TclTM.Exclude.ResourcePrefix
		}

		if aux.Modules[m].Exclude.Dependencies == nil {
			*aux.Modules[m].Exclude.Dependencies = *aux.TclTM.Exclude.Dependencies
		}

		if aux.Modules[m].Exclude.Provide == nil {
			*aux.Modules[m].Exclude.Provide = *aux.TclTM.Exclude.Provide
		}

		if aux.Modules[m].Exclude.SatisfyTcl == nil {
			*aux.Modules[m].Exclude.SatisfyTcl = *aux.TclTM.Exclude.SatisfyTcl
		}

		// Output
		if aux.Modules[m].Output == nil {
			aux.Modules[m].Output = aux.TclTM.Output
		}

		if aux.Modules[m].Output.Repository == nil {
			*aux.Modules[m].Output.Repository = *aux.TclTM.Output.Repository
		}

		if aux.Modules[m].Output.InteractiveLoader == nil {
			*aux.Modules[m].Output.InteractiveLoader = *aux.TclTM.Output.InteractiveLoader
		}
	}

	// Populate Configuration
	c.Version = aux.Version
	c.TclTM = aux.TclTM
	c.Modules = aux.Modules

	return nil
}

func (c *Config) String() string {
	b, _ := yaml.Marshal(c)
	return string(b)
}

// Resolve will resolve all variables within the configuration to their
// matching environment variables
//
// When strict mode is `false` then any error resolving variables
// will only result in a warning log message.
// When strict mode is `true` any resolving error will be treated as as error.
func (c *Config) Resolve(strict bool) (err error) {
	// Parse Modules
	for m := range c.Modules {
		// Resolve ENV in Name
		if c.Modules[m].Name, err = interpolate.Interpolate(Environment, c.Modules[m].Name); err != nil {
			r := log.WithError(err).WithFields(log.Fields{
				"Module":   m,
				"Property": "name",
				"Value":    c.Modules[m].Name,
			})

			if strict {
				r.Error(ErrResolveEnvFailed)
				return err
			}
			r.Warning(ErrResolveEnvFailed)
		}

		// Resolve ENV in Version
		if c.Modules[m].Version, err = interpolate.Interpolate(Environment, c.Modules[m].Version); err != nil {
			r := log.WithError(err).WithFields(log.Fields{
				"Module":   m,
				"Property": "version",
				"Value":    c.Modules[m].Version,
			})

			if strict {
				r.Error(ErrResolveEnvFailed)
				return err
			}
			r.Warning(ErrResolveEnvFailed)
		}

		// Resolve ENV in License
		// This will allow to use ${PATH} variables in 'license' key
		if c.Modules[m].License, err = interpolate.Interpolate(Environment, c.Modules[m].License); err != nil {
			r := log.WithError(err).WithFields(log.Fields{
				"Module":   m,
				"Property": "license",
				"Value":    c.Modules[m].License,
			})

			if strict {
				r.Error(ErrResolveEnvFailed)
				return err
			}
			r.Warning(ErrResolveEnvFailed)
		}

		// Resolve ENV in Extension
		if c.Modules[m].Extension, err = interpolate.Interpolate(Environment, c.Modules[m].Extension); err != nil {
			r := log.WithError(err).WithFields(log.Fields{
				"Module":   m,
				"Property": "extension",
				"Value":    c.Modules[m].Extension,
			})

			if strict {
				r.Error(ErrResolveEnvFailed)
				return err
			}
			r.Warning(ErrResolveEnvFailed)
		}

		// Resolve ENV in FinalName
		if c.Modules[m].FinalName, err = interpolate.Interpolate(Environment, c.Modules[m].FinalName); err != nil {
			r := log.WithError(err).WithFields(log.Fields{
				"Module":   m,
				"Property": "finalname",
				"Value":    c.Modules[m].FinalName,
			})

			if strict {
				r.Error(ErrResolveEnvFailed)
				return err
			}
			r.Warning(ErrResolveEnvFailed)
		}

		// Resolve ENV in meta
		for k, v := range c.Modules[m].Meta {
			val, err := interpolate.Interpolate(Environment, v)
			if err != nil {
				r := log.WithError(err).WithFields(log.Fields{
					"Module":   m,
					"Property": fmt.Sprintf("meta.%s", k),
					"Value":    v,
				})

				if strict {
					r.Error(ErrResolveEnvFailed)
					return err
				}
				r.Warning(ErrResolveEnvFailed)
			}
			c.Modules[m].Meta[k] = val
		}

		// Resolve ENV in filter
		for k, v := range c.Modules[m].Filter {
			val, err := interpolate.Interpolate(Environment, v)
			if err != nil {
				r := log.WithError(err).WithFields(log.Fields{
					"Module":   m,
					"Property": fmt.Sprintf("filter.%s", k),
					"Value":    v,
				})

				if strict {
					r.Error(ErrResolveEnvFailed)
					return err
				}
				r.Warning(ErrResolveEnvFailed)
			}
			c.Modules[m].Filter[k] = val
		}

		// Resolve ENV in files
		for f := range c.Modules[m].Files {
			// .name
			if c.Modules[m].Files[f].Name, err = interpolate.Interpolate(Environment, c.Modules[m].Files[f].Name); err != nil {
				r := log.WithError(err).WithFields(log.Fields{
					"Module":   m,
					"File":     f,
					"Property": "file.name",
					"Value":    c.Modules[m].Files[f].Name,
				})

				if strict {
					r.Error(ErrResolveEnvFailed)
					return err
				}
				r.Warning(ErrResolveEnvFailed)
			}

			// .filter
			for k, v := range c.Modules[m].Files[f].Filter {
				val, err := interpolate.Interpolate(Environment, v)
				if err != nil {
					r := log.WithError(err).WithFields(log.Fields{
						"Module":   m,
						"File":     f,
						"Property": fmt.Sprintf("file.filter.%s", k),
						"Value":    v,
					})

					if strict {
						r.Error(ErrResolveEnvFailed)
						return err
					}
					r.Warning(ErrResolveEnvFailed)
				}
				c.Modules[m].Files[f].Filter[k] = val
			}
		}

		// TODO: Resolve version from pkgIndex.tcl
	}

	return nil
}

// Validate the configuration
//
// When strict mode is `false` any error is returned as warnings.
// When strict mode is `true` any error is returned as an error.
func (c *Config) Validate(strict bool) (err error) {
	for i, m := range c.Modules {
		if len(m.Name) <= 0 {
			r := log.WithError(ErrPropertyRequired).WithFields(log.Fields{
				"Module":   i,
				"Property": "name",
			})

			if strict {
				r.Error(ErrModuleNameRequired)
				return ErrPropertyRequired
			}
			r.Warning(ErrModuleNameRequired)
		}

		if len(m.Version) <= 0 {
			r := log.WithError(ErrPropertyRequired).WithFields(log.Fields{
				"Module":   i,
				"Property": "version",
			})

			if strict {
				r.Error(ErrModuleVersionRequired)
				return ErrPropertyRequired
			}
			r.Warning(ErrModuleVersionRequired)
		}

		if len(m.Tcl) <= 0 {
			r := log.WithError(ErrPropertyRequired).WithFields(log.Fields{
				"Module":   i,
				"Property": "tcl",
			})

			if strict {
				r.Error(ErrModuleTclVersionRequired)
				return ErrPropertyRequired
			}
			r.Warning(ErrModuleTclVersionRequired)
		}
	}

	return nil
}

// EOF
