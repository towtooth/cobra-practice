package cmd

import (
	"errors"

	"gitlab.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s",
	short: "An encapsulation of kubectl",
	Long:  `An encapsulation of kubectl,but i'm long`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unregcognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
