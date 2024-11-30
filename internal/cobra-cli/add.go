package cobra_cli

import (
	"fmt"

	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:        "add",
	Aliases:    []string{"addition", "insert"},
	SuggestFor: []string{"ass", "ad"},
	Short:      "adds an task to your To-Do List",
	Long:       "The add method should be used to create new tasks in the underlying data store. It should take a positional argument with the task description",
	Example:    "tasks add <description>",
	Args:       cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addId := csvUtil.AddCsv(args[0])
		fmt.Printf("Task Added To Your To-Do List with ID: %s\n", addId)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
