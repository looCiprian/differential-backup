package command

import (
	"diff-backup/internal/db_mng"
	"diff-backup/internal/file_mng"
	"errors"
	"fmt"
)

type initCommand struct {
	destination string
}

//executeInit
// Execute directory initialization
func executeInit(initcommand initCommand) error {
	destination := file_mng.AddSlashIfNotPresent(initcommand.destination)

	if !file_mng.IsEmptyDirectory(destination) {
		fmt.Println("The directory is not empty ")
		for i, file := range file_mng.FilesInDirectory(destination){
			fmt.Println("File" + string(i) + " Name: " + file)
		}
		return errors.New("")
	}

	DBPath := destination + "index.db"

	if file_mng.CreateNewFileWithContent(destination + "IMPORTANT.txt", "DO NOT DELETE / ADD FILES MANUALLY, IF YOU NEED SOME DATA ONLY COPY OPERATIONS ARE ALLOWED") != nil{
		return errors.New("Cannot write important file ")
	}

	dbAlreadyExists := file_mng.FileExists(DBPath)
	if !dbAlreadyExists {
		err := db_mng.OpenDB(DBPath)
		if err != nil {
			return err
		}
		db_mng.CreateTable()
		db_mng.CloseDB()
	} else {
		return errors.New("Backup directory already in use ")
	}

	fmt.Println("Backup correctly initialized ")

	return nil
}


