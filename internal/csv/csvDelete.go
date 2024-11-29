package csv

func DeleteTask(id string) bool {
	tasks := ReadCsv()

	updatedTasks := []*task{}

	change := false

	for _, task := range tasks {
		if task.Id != id {
			updatedTasks = append(updatedTasks, task)
			change = true
		}
	}

	WriteCsv(updatedTasks)

	return change

}