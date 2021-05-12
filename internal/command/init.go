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

func executeInit(initcommand initCommand) error {
	destination := initcommand.destination

	if !file_mng.IsEmptyDirectory(destination){
		return errors.New("The directory is not empty ")
	}

	destination = file_mng.AddSlashIfNotPresent(destination) + "index.db"

	dbAlreadyExists := file_mng.FileExists(destination)
	if !dbAlreadyExists {
		database,_ := db_mng.OpenDB(destination)
		db_mng.CreateTable(*database)
	}else {
		return errors.New("Backup directory already in use ")
	}

	fmt.Println("Backup correctly initialized ")

	return nil
}