package cobra_cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A cli application for managing tasks in the terminal.",
	Long:  "A cli application for managing tasks in the terminal. Should be able to perform crud operations via a cli on a data file of tasks. The operations should be as follows: [add, list, complete, delete]",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured: %s\n", err)
		os.Exit(1)
	}
}
