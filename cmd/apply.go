package cmd

import (
	"errors"

	"github.com/buddyspike/mbt/lib"
	"github.com/spf13/cobra"
)

var to, out string

func init() {
	ApplyCmd.PersistentFlags().StringVar(&to, "to", "", "template to apply")
	ApplyCmd.PersistentFlags().StringVar(&out, "out", "", "output path")
	ApplyCmd.AddCommand(ApplyBranchCmd)
	RootCmd.AddCommand(ApplyCmd)
}

var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "applies a manifest over a template",
}

var ApplyBranchCmd = &cobra.Command{
	Use:   "branch <branch>",
	Short: "applies the manifest of specified branch over a given template",
	RunE: func(cmd *cobra.Command, args []string) error {
		out := ""
		branch := "master"
		if len(args) > 0 {
			branch = args[0]
		}

		if to == "" {
			return errors.New("requires the path to template")
		}

		return lib.ApplyBranch(globalArgIn, to, branch, out)
	},
}