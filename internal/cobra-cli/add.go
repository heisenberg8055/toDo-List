package cobra_cli

import (
	"fmt"

	sqlite_server "github.com/heisenberg8055/toDo-List/internal/sqlite-server"
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
		chk, err := sqlite_server.AddTask(args[0])
		if err != nil {
			panic(err)
		}
		if chk != 0 {
			fmt.Printf("Task Added To Your To-Do List with ID:%d\n", chk)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
