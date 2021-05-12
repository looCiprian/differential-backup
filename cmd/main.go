package main

import (
	"diff-backup/internal/command"
	"diff-backup/internal/menu"
	"fmt"
)

func main() {
	for {
		commandToExecute := menu.Base("Press TAB for options")
		error := command.Execute(commandToExecute)
		if error != nil {
			fmt.Println(error)
		}
	}
	//command.ExecuteBackupTest()
}
