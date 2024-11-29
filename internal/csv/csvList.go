package csv

func ListCsv(flag string) []*task {

	tasks := ReadCsv()
	if flag == "" {
		notCompletedTasks := []*task{}
		for _, task := range tasks {
			if !task.IsComplete {
				notCompletedTasks = append(notCompletedTasks, task)
			}
		}
		return notCompletedTasks
	} else {
		return tasks
	}

}
