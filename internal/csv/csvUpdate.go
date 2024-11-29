package csv

func CompleteTask(id string) bool {
	tasks := ReadCsv()

	updatedTasks := []*task{}

	change := false

	for _, task := range tasks {
		if task.Id == id {
			task.IsComplete = !task.IsComplete
			change = true
		}
		updatedTasks = append(updatedTasks, task)
	}

	WriteCsv(updatedTasks)

	return change

}
