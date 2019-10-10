package commands

import (
	"fmt"
	"strings"

	"github.com/gjrtimmer/tcltm/pkg/config"
	"github.com/gjrtimmer/viper"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	// Verbose Logging
	Verbose = false

	// OptionInputDirectory ..
	OptionInputDirectory string

	// OptionOutputDirectory ...
	OptionOutputDirectory string

	// OptionConfigFile ..
	OptionConfigFile string

	// OptionPackageName ...
	OptionPackageName string

	// Config struct
	Config = config.Config{}
)

var (
	// Version ...
	Version = "latest"

	// LoadConfig is the generic PrePersistent function to load
	// the viper config
	LoadConfig = func(cmd *cobra.Command, args []string) (err error) {
		viper.AddConfigPath(viper.GetString("tcltm.input"))
		if err = viper.ReadInConfig(); err != nil {
			return err
		}

		if err = viper.Unmarshal(&Config); err != nil {
			return err
		}

		if Verbose {
			var raw []byte
			if raw, err = yaml.Marshal(&Config); err != nil {
				return err
			}

			fmt.Printf("%s\nConfiguration:\n %s%s\n", strings.Repeat("-", 80), raw, strings.Repeat("-", 80))
		}

		return nil
	}

	// MainCommand ...
	MainCommand = &cobra.Command{
		Use:     "tcltm",
		Short:   "Tcl Module Builder",
		Long:    ``,
		Version: Version,
	}
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TCLTM")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".tcltm")

	// Command Configuration
	MainCommand.SetHelpTemplate(helpTemplate())

	// Flag Defaults
	viper.SetDefault("tcltm.verbose", false)
	viper.SetDefault("tcltm.input", ".")
	viper.SetDefault("tcltm.output", ".")
	viper.SetDefault("tcltm.config", ".tcltm")
	viper.SetDefault("pkg", "")

	// Global Flags
	MainCommand.PersistentFlags().BoolVar(&Verbose, "verbose", false, "Show verbose logging")
	MainCommand.PersistentFlags().StringVarP(&OptionInputDirectory, "in", "i", ".", "Input directory")
	MainCommand.PersistentFlags().StringVarP(&OptionOutputDirectory, "out", "o", ".", "Output directory")
	MainCommand.PersistentFlags().StringVarP(&OptionConfigFile, "config", "c", ".tcltm", "Alternate config file")
	MainCommand.PersistentFlags().StringVarP(&OptionPackageName, "pkg", "p", "", "Only build package name (default build all)")

	// Environment Binding
	viper.BindPFlag("tcltm.verbose", MainCommand.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("tcltm.input", MainCommand.PersistentFlags().Lookup("in"))
	viper.BindPFlag("tcltm.output", MainCommand.PersistentFlags().Lookup("out"))
	viper.BindPFlag("tcltm.config", MainCommand.PersistentFlags().Lookup("config"))
	viper.BindPFlag("pkg", MainCommand.PersistentFlags().Lookup("pkg"))
}

func helpTemplate() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsagesWrapped 80 | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsagesWrapped 80 | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}

// EOF
