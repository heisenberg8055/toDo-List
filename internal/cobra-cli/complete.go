package cobra_cli

import (
	"fmt"

	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:        "complete",
	Aliases:    []string{"flip", "mark"},
	SuggestFor: []string{"compete", "done"},
	Short:      "Change task Status as True or False",
	Long:       "Flip the status of the tasks in the To-Do List",
	Args:       cobra.ExactArgs(1),
	Example:    "tasks complete <taskid>",
	Run: func(cmd *cobra.Command, args []string) {
		chk := csvUtil.CompleteTask(args[0])
		if chk {
			fmt.Printf("Status Flipped for ID: %s\n", args[0])
		} else {
			fmt.Printf("Wrong ID Passed: %s, use tasks list [-a, --all] to check for a valid ID.\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
