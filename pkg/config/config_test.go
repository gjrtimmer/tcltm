package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.timmertech.nl/go/interpolate"
)

func TestMain(m *testing.M) {
	Environment = interpolate.NewSliceEnv([]string{
		"NAME=env-resolve",
		"VERSION=1.0.0",
		"PATH=/tmp",
		"EXT=tmp",
		"USER=env-test",
		"OS=linux",
	})

	os.Exit(m.Run())
}

func TestLoad(t *testing.T) {
	c, err := Load("test/tcltm.yml", false)
	if err != nil {
		t.Fail()
	}

	// TODO: Assertions
	// Config.TclTm.Version
	assert.Equal(t, "4.0", c.Version, "ShouldBeEqual")

	// Config.TclTm.Include
	assert.Equal(t, true, *c.TclTM.Include.Require, "ShouldBeEqual")

	// Config.TclTm.Exclude
	assert.Equal(t, true, *c.TclTM.Exclude.Comments, "ShouldBeEqual")
	assert.Equal(t, true, *c.TclTM.Exclude.ResourcePrefix, "ShouldBeEqual")
	assert.Equal(t, true, *c.TclTM.Exclude.Dependencies, "ShouldBeEqual")
	assert.Equal(t, true, *c.TclTM.Exclude.Provide, "ShouldBeEqual")
	assert.Equal(t, true, *c.TclTM.Exclude.SatisfyTcl, "ShouldBeEqual")

	// Config.TclTm.Output
	assert.Equal(t, true, *c.TclTM.Output.Repository, "ShouldBeEqual")
	assert.Equal(t, true, *c.TclTM.Output.InteractiveLoader, "ShouldBeEqual")

	assert.Len(t, c.Modules, 1)
	assert.Equal(t, "config-test", c.Modules[0].Name)
	assert.Equal(t, "0.0.0", c.Modules[0].Version)
	assert.Equal(t, "8.6", c.Modules[0].Tcl)
	assert.Equal(t, "tclsh", c.Modules[0].Interpreter)
	assert.Equal(t, "Config Test", c.Modules[0].Summary)
	assert.NotEmpty(t, c.Modules[0].Description)
	assert.Contains(t, c.Modules[0].Description, "\n")
	assert.Equal(t, "LICENSE", c.Modules[0].License)
	assert.Len(t, c.Modules[0].Dependencies, 2)
	assert.Equal(t, "tm", c.Modules[0].Extension)
	assert.Equal(t, "config-test.tm", c.Modules[0].FinalName)

	assert.Len(t, c.Modules[0].Meta, 1)
	assert.Equal(t, "bar", c.Modules[0].Meta["foo"])

	assert.Len(t, c.Modules[0].Filter, 1)
	assert.Equal(t, "test", c.Modules[0].Filter["user"])

	assert.Len(t, c.Modules[0].Files, 1)
	assert.Equal(t, "test.tcl", c.Modules[0].Files[0].Name)
	assert.Equal(t, FileTypeBinary, c.Modules[0].Files[0].Type)
	assert.Equal(t, FileActionRun, c.Modules[0].Files[0].Action)
	assert.Equal(t, true, c.Modules[0].Files[0].Filtering)
	assert.Len(t, c.Modules[0].Files[0].Filter, 1)
	assert.Equal(t, "linux", c.Modules[0].Files[0].Filter["os"])
}

func TestResolveEnv(t *testing.T) {
	c, err := Load("test/env.yml", false)
	if err != nil {
		t.Fail()
	}

	// TODO: Assertions
	assert.Equal(t, "env-resolve", c.Modules[0].Name)
	assert.Equal(t, "1.0.0", c.Modules[0].Version)
	assert.Equal(t, "/tmp/LICENSE", c.Modules[0].License)
	assert.Equal(t, "env-test", c.Modules[0].Meta["username"])
	assert.Equal(t, "env-test", c.Modules[0].Filter["user"])
	assert.Equal(t, "tmp", c.Modules[0].Extension)
	assert.Equal(t, "env-resolve-1.0.0.tmp", c.Modules[0].FinalName)
	assert.Equal(t, "linux", c.Modules[0].Files[0].Filter["os"])
}

func TestOverride(t *testing.T) {
	_, err := Load("test/override.yml", false)
	if err != nil {
		t.Fail()
	}

	// TODO: Assertions
}

// EOF
