package cobra_cli

import (
	"fmt"

	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"flip", "mark"},
	Short:   "Change task Status",
	Long:    "Flip the status of the tasks in the To-Do List",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		chk := csvUtil.CompleteTask(args[0])
		if chk {
			fmt.Printf("Status Flipped for ID: %s", args[0])
		} else {
			fmt.Printf("Wrong ID Passed: %s", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
