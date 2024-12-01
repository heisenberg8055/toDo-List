package sqlite_server

func ListTask(showAll bool) {
	tasks := ReadTasks(db)
	if !showAll {
		notCompleted := []task{}
		for _, task := range tasks {
			if !task.IsComplete {
				notCompleted = append(notCompleted, task)
			}
		}
		StdOut(notCompleted)
	} else {
		StdOut(tasks)
	}
}
