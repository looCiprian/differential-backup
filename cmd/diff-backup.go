package main

import (
	"diff-backup/internal/command"
	"diff-backup/internal/menu"
	"fmt"
)

func main() {
	for {
		commandToExecute := menu.Base("Press TAB for options")
		err := command.Execute(commandToExecute)
		if err != nil {
			fmt.Println(err)
		}
	}
}
