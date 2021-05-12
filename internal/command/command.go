package command

import (
	"diff-backup/internal/menu"
	"errors"
	"fmt"
	"strings"
)

func Execute(command string) error {

	if len(strings.Fields(command)) == 0 {
		return errors.New("")
	}

	switch strings.Fields(command)[0]{
	case "init":
		initcommand := initCommandParser(command)
		error := executeInit(initcommand)
		if error != nil{
			return error
		}
	case "backup":
		backupcommand, error := backupCommandParser(command)
		if error != nil{
			return error
		}
		executeBackup(backupcommand)
	case "restore":
		fmt.Printf("Not implemented restore")
	case "exit":
		executeExit()
	default:
		menu.Base("Wrong command. Press TAB for help")
	}
	return nil
}

func backupCommandParser(command string) (backupCommand, error) {
	if len(strings.Fields(command)) != 5 || !strings.Contains(command,"--source") || !strings.Contains(command,"--destination"){
		return backupCommand{}, errors.New("Wrong backup parameters")
	}
	var backupCommand = backupCommand{}

	if strings.Fields(command)[1] == "--destination" && strings.Fields(command)[3] == "--source"{
		backupCommand.destination = strings.Fields(command)[2]
		backupCommand.source = strings.Fields(command)[4]
	}

	if strings.Fields(command)[3] == "--destination" && strings.Fields(command)[1] == "--source"{
		backupCommand.destination = strings.Fields(command)[4]
		backupCommand.source = strings.Fields(command)[2]
	}
	return backupCommand, nil
}

func initCommandParser(command string) initCommand {
	if len(strings.Fields(command)) != 3{
		menu.Base("Wrong init parameters")
	}

	var initcommand = initCommand{}

	if strings.Fields(command)[1] == "--destination"{
		initcommand.destination = strings.Fields(command)[2]
	}

	return initcommand
}