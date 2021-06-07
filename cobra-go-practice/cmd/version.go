package cmd

import (
	"fmt"
	"os"

	"gitlab.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show k8s version info.",

	Run: func(cmd cobra.Command, args []string) {
		output, err = ExecuteCommand("kubectl", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.addCommand(versionCmd)
}
