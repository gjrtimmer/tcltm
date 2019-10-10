package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	// GlobalFlags defintion
	GlobalFlags []cli.Flag
)

func init() {
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Default only log warning or above
	log.SetLevel(log.WarnLevel)

	// Register Caller
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	// Configure Global Flags
	GlobalFlags = []cli.Flag{
		cli.BoolFlag{
			Name:   "verbose",
			EnvVar: "TCLTM_DEBUG",
		},
	}
}

// EOF
