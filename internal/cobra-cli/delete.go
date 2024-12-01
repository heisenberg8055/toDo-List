package cobra_cli

import (
	"fmt"

	sqlite "github.com/heisenberg8055/toDo-List/internal/sqlite-server"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:        "delete",
	Aliases:    []string{"remove", "yeet"},
	SuggestFor: []string{"del", "fuck", "rem"},
	Short:      "Delete a task from the data store",
	Long:       "Remove the record from the To-Do list based on task ID",
	Example:    "tasks delete <taskid>",
	Args:       cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		chk := sqlite.DeleteTask(args[0])
		if chk {
			fmt.Printf("Removed record with ID: %s\n", args[0])
		} else {
			fmt.Printf("Wrong ID Passed: %s, use tasks list [-a, -all] to check for a valid ID.\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
