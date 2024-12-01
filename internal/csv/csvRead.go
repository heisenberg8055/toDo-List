package csv

import (
	"errors"
	"os"

	"github.com/gocarina/gocsv"
)

func ReadCsv() []*task {

	tasks := []*task{}
	if _, err := os.Stat("./assets/tasks.csv"); errors.Is(err, os.ErrNotExist) {
		return tasks
	}
	taskFile, err := LoadFile("./assets/tasks.csv")
	defer CloseFile(taskFile)

	if err != nil {
		panic(err)
	}

	if _, err := taskFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	if err := gocsv.UnmarshalFile(taskFile, &tasks); err != nil {
		if err == gocsv.ErrEmptyCSVFile {
			return []*task{}
		}
	}

	return tasks
}
