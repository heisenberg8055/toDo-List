package main

import (
	cliUtil "github.com/heisenberg8055/toDo-List/internal/cobra-cli"
	dbUtil "github.com/heisenberg8055/toDo-List/internal/sqlite-server"
)

func main() {
	dbUtil.InitDatabase("./assets/tasks.db")
	cliUtil.Execute()
}
