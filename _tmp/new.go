package commands

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/gjrtimmer/viper"
	"github.com/spf13/cobra"
)

var (
	// NewCommand ...
	NewCommand = &cobra.Command{
		Use:   "new",
		Short: "Create new ...",
		Long:  ``,
	}

	// NewConfigCommand will create a new configuration file
	// at either the provided output directory or oif provided
	// the config file location.
	NewConfigCommand = &cobra.Command{
		Use:   "config",
		Short: "Create new config",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := viper.GetString("tcltm.config")
			path := filepath.Join(viper.GetString("tcltm.input"), cfg)

			if Verbose {
				fmt.Printf("writing new config: %s\n", path)
			}

			var err error
			var raw []byte
			if raw, err = yaml.Marshal(&Config); err != nil {
				return fmt.Errorf("failed to marshall config: %s", err)
			}
			if err = ioutil.WriteFile(path, raw, 0775); err != nil {
				return fmt.Errorf("failed to write config: %s", err)
			}

			return nil
		},
	}
)

func init() {
	NewCommand.AddCommand(NewConfigCommand)

	// To to MainCommand
	MainCommand.AddCommand(NewCommand)
}

// EOF
