package sqlite_server

func DeleteTask(taskID string) bool {
	return DeleteTaskInDB(db, taskID)
}
