package csv

func ListCsv(showAll bool) {

	tasks := ReadCsv()
	if !showAll {
		notCompletedTasks := []*task{}
		for _, task := range tasks {
			if !task.IsComplete {
				notCompletedTasks = append(notCompletedTasks, task)
			}
		}
		StdOut(notCompletedTasks)
	} else {
		StdOut(tasks)
	}
}
