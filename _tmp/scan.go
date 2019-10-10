package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gjrtimmer/tcltm/pkg/osutil"
	"github.com/gjrtimmer/viper"
	"github.com/spf13/cobra"
)

var (
	// Scanning RegExp
	//DependencyRegExp = regexp.MustCompile("package (provide|require|ifneeded)(?:[[:blank:]]+)([_[[:alpha:]]][:_[[:alnum:]]]*)(?:])?((?:[[:blank:]]+)?(?:([\\d]+.)?([\\d]+.)?(\\*|[\\d]+))?)")
	DependencyRegExp = regexp.MustCompile(`package (provide|require|ifneeded)(?:[[:blank:]]+)([_[:alpha:]][:_[:alnum:]]*)(?:])?((?:[[:blank:]]+)?(?:(\d+\.)?(\d+\.)?(\*|\d+))?)`)
	// ([_[:alpha:]][:_[:alnum:]]*)(?:\])?((?:[[:blank:]]+)?(?:(\d+\.)?(\d+\.)?(\*|\d+))?)
	// ScanCommand defines the global command
	// for scanning operations
	ScanCommand = &cobra.Command{
		Use:   "scan",
		Short: "Scan operations",
		Long:  ``,
	}

	// DependencyScanCommand ...
	DependencyScanCommand = &cobra.Command{
		Use:   "deps",
		Short: "Scan for Tcl dependencies",
		Long: `Scan for Tcl dependencies.

If the scan location is a directory all *.tcl files will
be scanned for dependencies.`,
		Aliases: []string{"dependencies", "dependency"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if Verbose {
				fmt.Printf("Scanning: %s\n", viper.GetString("tcltm.input"))
			}

			_, err := scan(viper.GetString("tcltm.input"))
			if err != nil {
				return fmt.Errorf("scanner: %s", err)
			}

			return nil
		},
	}
)

func init() {
	ScanCommand.AddCommand(DependencyScanCommand)
	MainCommand.AddCommand(ScanCommand)
}

type tclPackage struct {
	Name    string
	Version string
	Type    string
}

func scan(path string) ([]tclPackage, error) {
	var pkg []tclPackage
	var err error

	if ok, err := osutil.Exists(path); !ok {
		return nil, fmt.Errorf("path does not exists: %s", err)
	}

	// Get Abolsute Path
	if path, err = osutil.Abs(path); err != nil {
		return nil, err
	}
	if ok, err := osutil.IsDirectory(path); !ok {
		if path, err = osutil.Dir(path); err != nil {
			return nil, err
		}
	}

	// Scan for Tcl files
	files, err := osutil.List(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if filepath.Ext(f) == ".tcl" || filepath.Ext(f) == ".tm" {
			file, err := os.Open(f)
			if err != nil {
				return nil, err
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Bytes()
				if bytes.ContainsRune(line, '\u001A') {
					break
				}

				res := DependencyRegExp.FindAllSubmatch(line, -1)
				fmt.Printf("%+s", res)
			}

			// Close file
			file.Close()
		}
	}

	return pkg, nil
}

// EOF
