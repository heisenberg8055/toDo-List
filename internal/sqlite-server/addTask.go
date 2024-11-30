package sqlite_server

import (
	"log"

	timeUtil "github.com/heisenberg8055/toDo-List/internal/timediff"
)

func AddTask(addDescription string) (int64, error) {
	addTask := task{Id: 0, Description: addDescription, CreatedAt: timeUtil.GetCurrentTime(), IsComplete: false}
	db, err := InitDatabase("./assets/tasks.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return 0, err
	}
	return AddTaskToDb(db, addTask)
}
