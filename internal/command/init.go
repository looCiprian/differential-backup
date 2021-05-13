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
	destination := initcommand.destination

	if !file_mng.IsEmptyDirectory(destination) {
		return errors.New("The directory is not empty ")
	}

	destination = file_mng.AddSlashIfNotPresent(destination) + "index.db"

	dbAlreadyExists := file_mng.FileExists(destination)
	if !dbAlreadyExists {
		database, err := db_mng.OpenDB(destination)
		if err != nil {
			return err
		}
		db_mng.CreateTable(*database)
		db_mng.CloseDB(*database)
	} else {
		return errors.New("Backup directory already in use ")
	}

	fmt.Println("Backup correctly initialized ")

	return nil
}
