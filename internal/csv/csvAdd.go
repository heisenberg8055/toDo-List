package csv

import (
	"strconv"

	timeUtil "github.com/heisenberg8055/toDo-List/internal/timediff"
)

func AddCsv(addDescription string) string {

	tasks := ReadCsv()

	addCreatedAt := timeUtil.GetCurrentTime()

	if len(tasks) == 0 {
		tasks = append(tasks, &task{Id: "1", Description: addDescription, CreatedAt: addCreatedAt, IsComplete: false})
		WriteCsv(tasks)
		return "1"
	} else {
		lastId, err := strconv.Atoi(tasks[len(tasks)-1].Id)
		if err != nil {
			panic(err)
		}
		lastId += 1
		addId := strconv.Itoa(lastId)
		tasks = append(tasks, &task{Id: addId, Description: addDescription, CreatedAt: addCreatedAt, IsComplete: false})
		WriteCsv(tasks)
		return addId
	}

}
