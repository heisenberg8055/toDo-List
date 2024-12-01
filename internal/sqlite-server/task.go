package sqlite_server

import (
	"database/sql"
	"log"
)

type task struct {
	Id          int
	Description string
	CreatedAt   string
	IsComplete  bool
}

func AddTaskToDb(db *sql.DB, newTask task) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO TASKS (id, description, createdAt, isComplete) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	res, _ := stmt.Exec(nil, newTask.Description, newTask.CreatedAt, newTask.IsComplete)
	defer stmt.Close()
	return res.LastInsertId()
}

func ReadTasks(db *sql.DB) []task {
	rows, err := db.Query("SELECT * FROM TASKS")
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tasks := make([]task, 0)
	for rows.Next() {
		currTask := task{}
		err = rows.Scan(&currTask.Id, &currTask.Description, &currTask.CreatedAt, &currTask.IsComplete)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, currTask)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tasks
}

func GetTaskById(db *sql.DB, taskID string) task {
	rows, _ := db.Query("SELECT * from tasks where id = '" + taskID + "'")
	defer rows.Close()
	currTask := task{}
	for rows.Next() {
		rows.Scan(&currTask.Id, &currTask.Description, &currTask.CreatedAt, &currTask.IsComplete)
	}
	return currTask
}

func UpdateTaskInDB(db *sql.DB, taskID string) bool {

	currTask := GetTaskById(db, taskID)

	if (task{}) == currTask {
		return false
	}

	chk := currTask.IsComplete

	chk = !chk

	stmt, err := db.Prepare("UPDATE tasks set isComplete = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(chk, taskID)
	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteTaskInDB(db *sql.DB, taskID string) bool {

	currTask := GetTaskById(db, taskID)

	if (task{}) == currTask {
		return false
	}

	stmt, err := db.Prepare("DELETE FROM tasks where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(taskID)
	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return affected != 0
}
