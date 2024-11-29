package cobra_cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "a cli To-Do list",
	Long:  "Manage your To-Do List in your own Cli",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured: %s\n", err)
		os.Exit(1)
	}
}
