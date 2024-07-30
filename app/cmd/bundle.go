package cmd

import (
	"github.com/spf13/cobra"
)

func NewBundleCmd() *cobra.Command {
	bundleCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	bundleCmd.AddCommand(NewInsertCmd())
	return bundleCmd
}
