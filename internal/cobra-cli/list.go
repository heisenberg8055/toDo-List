package cobra_cli

import (
	csvUtil "github.com/heisenberg8055/toDo-List/internal/csv"
	"github.com/spf13/cobra"
)

var showAll bool
var listCmd = &cobra.Command{
	Use:        "list",
	Aliases:    []string{"show", "flash"},
	SuggestFor: []string{"lis", "liat"},
	Short:      "Return a list of all of the uncompleted tasks",
	Long:       "Return a list of all of the uncompleted tasks, with the option to return all tasks regardless of whether or not they are completed.",
	Example:    "tasks list <optional> (-a or --all)",
	Run: func(cmd *cobra.Command, args []string) {
		csvUtil.ListCsv(showAll)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Shows all Tasks")
	rootCmd.AddCommand(listCmd)
}
