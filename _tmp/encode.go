package commands

import (
	"os"

	"github.com/gjrtimmer/viper"
	"github.com/spf13/cobra"
)

var (
	// EncodeCommand ...
	EncodeCommand = &cobra.Command{
		Use:               "encode",
		Short:             "Encode Tcl Module",
		Long:              ``,
		PersistentPreRunE: LoadConfig,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
				os.Exit(0)
			}
		},
	}
)

func init() {
	// Flags Defaults
	viper.SetDefault("tcltm.strip.comments", false)
	viper.SetDefault("tcltm.strip.resourcedir", false)
	viper.SetDefault("tcltm.exclude.satisfy-tcl", false)
	viper.SetDefault("tcltm.exclude.deps", false)
	viper.SetDefault("tcltm.exclude.provide", false)
	viper.SetDefault("tcltm.preserve.require", false)
	viper.SetDefault("tcltm.interactive-loader", false)
	viper.SetDefault("tcltm.repository", false)

	// Local Flags
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Strip.Comments, "strip-comments", false, "Strip comments from source")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Strip.ResourceDir, "strip-resource-dir", false, "Strip the directory from the source file path")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Exclude.SatisfyTcl, "exclude-satisfy-tcl", false, "Exclude Tcl vsatisfies command")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Exclude.Dependencies, "exclude-deps", false, "Exclude 'package require' commands for dependencies of generated module")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Exclude.Provide, "exclude-provide", false, "Exclude 'package provide' command")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Preserve.Require, "preserve-require", false, "Preserve 'package require' in source code")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.InteractiveLoader, "interactive-loader", false, "Enable interactive loader. Interactive loader will only run the binary loader when the tcl interpreter is in interactive mode")
	EncodeCommand.PersistentFlags().BoolVar(&Config.TclTM.Repository, "repository", false, "Create repository output directories.\n(tcl8/tcl<version>/<module>.tm")

	// Environment Binding
	viper.BindPFlag("tcltm.strip.comments", EncodeCommand.PersistentFlags().Lookup("strip-comments"))
	viper.BindPFlag("tcltm.strip.resourcedir", EncodeCommand.PersistentFlags().Lookup("strip-resource-dir"))
	viper.BindPFlag("tcltm.exclude.satisfy-tcl", EncodeCommand.PersistentFlags().Lookup("exclude-satisfy-tcl"))
	viper.BindPFlag("tcltm.exclude.deps", EncodeCommand.PersistentFlags().Lookup("exclude-deps"))
	viper.BindPFlag("tcltm.exclude.provide", EncodeCommand.PersistentFlags().Lookup("exclude-provide"))
	viper.BindPFlag("tcltm.preserve.require", EncodeCommand.PersistentFlags().Lookup("preserve-require"))
	viper.BindPFlag("tcltm.interactive-loader", EncodeCommand.PersistentFlags().Lookup("interactive-loader"))
	viper.BindPFlag("tcltm.repository", EncodeCommand.PersistentFlags().Lookup("repository"))

	// Add Command
	MainCommand.AddCommand(EncodeCommand)
}

// EOF
