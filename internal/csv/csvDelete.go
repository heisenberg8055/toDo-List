package csv

func DeleteTask(id string) bool {
	tasks := ReadCsv()

	updatedTasks := []*task{}

	for _, task := range tasks {
		if task.Id != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	WriteCsv(updatedTasks)

	return len(updatedTasks) != len(tasks)

}
