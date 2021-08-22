package command

/*
//Execute
// Send the user command to the correct function
func Execute(command string) error {

	if len(strings.Fields(command)) == 0 {
		return errors.New("")
	}

	switch strings.Fields(command)[0] {
	case "init":
		initcommand, err := initCommandParser(command)
		if err != nil {
			return err
		}
		err = executeInit(initcommand)
		if err != nil {
			return err
		}
	case "backup":
		backupcommand, err := backupCommandParser(command)
		if err != nil {
			return err
		}
		err = executeBackup(backupcommand)
		if err != nil {
			return err
		}
	case "restore":
		restorecommand, err := restoreCommandParser(command)
		if err != nil {
			return err
		}
		err = executeRestore(restorecommand)
		if err != nil {
			return err
		}
	case "exit":
		executeExit()
	default:
		menu.Base("Wrong command. Press TAB for help")
	}
	return nil
}

func restoreCommandParser(command string) (restoreCommand, error) {

	if len(strings.Fields(command)) != 7 ||
		!strings.Contains(command, "--source") ||
		!strings.Contains(command, "--destination") ||
		!strings.Contains(command, "--date") {
		return restoreCommand{}, errors.New("Wrong restore parameters ")
	}

	var restorecommand = restoreCommand{}
	commandArray := strings.Fields(command)

	i := 0
	for {

		if commandArray[i] == "--date" {
			restorecommand.date =  commandArray[i+1]
			i = i+1
		}
		if commandArray[i] == "--destination" {
			restorecommand.destination =  commandArray[i+1]
			i = i+1
		}
		if commandArray[i] == "--source" {
			restorecommand.source =  commandArray[i+1]
			i = i+1
		}
		i = i+1

		if i >= len(commandArray) {
			break
		}
	}


	return restorecommand, nil

}

//backupCommandParser
// Parse backup command
func backupCommandParser(command string) (backupCommand, error) {
	if len(strings.Fields(command)) != 5 || !strings.Contains(command, "--source") || !strings.Contains(command, "--destination") {
		return backupCommand{}, errors.New("Wrong backup parameters ")
	}
	var backupCommand = backupCommand{}

	if strings.Fields(command)[1] == "--destination" && strings.Fields(command)[3] == "--source" {
		backupCommand.destination = strings.Fields(command)[2]
		backupCommand.source = strings.Fields(command)[4]
	}

	if strings.Fields(command)[3] == "--destination" && strings.Fields(command)[1] == "--source" {
		backupCommand.destination = strings.Fields(command)[4]
		backupCommand.source = strings.Fields(command)[2]
	}
	return backupCommand, nil
}

//initCommandParser
// Parse init command
func initCommandParser(command string) (initCommand, error) {
	var initcommand = initCommand{}

	if len(strings.Fields(command)) != 3 {
		return initcommand, errors.New("Wrong init parameters ")
	}

	if strings.Fields(command)[1] == "--destination" {
		initcommand.destination = strings.Fields(command)[2]
	}else {
		return initcommand, errors.New("Wrong init parameters ")
	}

	return initcommand, nil
}
*/