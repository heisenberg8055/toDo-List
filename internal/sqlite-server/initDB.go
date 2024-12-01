package sqlite_server

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase(dbPath string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		_, err = db.Exec(
			`CREATE TABLE IF NOT EXISTS tasks (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				description TEXT NOT NULL, 
				createdAt TEXT NOT NULL, 
				isComplete BOOLEAN NOT NULL
			)`,
		)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
