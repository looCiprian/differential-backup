package main

import "diff-backup/cmd"

func main() {

	cmd.Execute()

	/*for {
		commandToExecute := menu.Base("Press TAB for options")
		err := command.Execute(commandToExecute)
		if err != nil {
			fmt.Println(err)
		}
	}*/

}