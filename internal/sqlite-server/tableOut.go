package sqlite_server

import (
	"os"

	timeUtil "github.com/heisenberg8055/toDo-List/internal/timediff"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func StdOut(tasks []task) {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Tasks")
	t.Style().Title.Align = text.AlignCenter
	t.AppendHeader(table.Row{"ID", "Task", "Created", "Done"})
	for _, task := range tasks {
		var chk = "X"
		if task.IsComplete {
			chk = "\u2713"
		}
		t.AppendRow([]interface{}{task.Id, task.Description, timeUtil.GetReadableTime(task.CreatedAt), chk})
	}
	t.Render()
}
