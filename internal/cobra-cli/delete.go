package cobra_cli

import (
	"fmt"

	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "yeet"},
	Short:   "Remove the record",
	Long:    "Remove the record from the To-Do list based on ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		chk := csvUtil.DeleteTask(args[0])
		if chk {
			fmt.Printf("Removed record with ID: %s\n", args[0])
		} else {
			fmt.Printf("Wrong ID Passed: %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
