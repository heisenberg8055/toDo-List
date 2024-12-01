package csv

import (
	"os"

	"github.com/gocarina/gocsv"
)

type task struct {
	Id          string `csv:"ID"`
	Description string `csv:"Description"`
	CreatedAt   string `csv:"CreatedAt"`
	IsComplete  bool   `csv:"IsComplete"`
}

func WriteCsv(tasks []*task) {

	taskFile, err := LoadFile("./assets/tasks.csv")
	defer CloseFile(taskFile)

	if err := os.Truncate("./assets/tasks.csv", 0); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	err = gocsv.MarshalFile(&tasks, taskFile)

	if err != nil {
		panic(err)
	}
}
