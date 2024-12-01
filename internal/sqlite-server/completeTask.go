package sqlite_server

func CompleteTask(taskID string) bool {
	return UpdateTaskInDB(db, taskID)
}
