package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Flags available to all commands.
var (
	in     string
	src    string
	dst    string
	from   string
	to     string
	first  string
	second string
	kind   string
	all    bool
)

func init() {
	RootCmd.PersistentFlags().StringVar(&in, "in", "", "path to repo")
}

// RootCmd is the main command.
var RootCmd = &cobra.Command{
	Use:   "mbt",
	Short: "Build utility for monorepos",
	Long: `Build utility for monorepos

Monorepo Build Tool (mbt) is a utility that supports differential builds,
dependency tracking and metadata management for monorepos stored in git.

See help for individual commands for more information.

	`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Use != "version" && in == "" {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			in = cwd
		}
		return nil
	},
}
