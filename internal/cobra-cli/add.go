package cobra_cli

import (
	"fmt"

	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition", "insert"},
	Short:   "adds an task to your To-Do List",
	Long:    "Adds a task to your To-Do List along with time and checklist",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		csvUtil.AddCsv(args[0])
		fmt.Printf("Task Added To Your To-Do List")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
