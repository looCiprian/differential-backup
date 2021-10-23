package command

import (
	"errors"
	"fmt"
	"github.com/looCiprian/diff-backup/internal/config"
	"github.com/looCiprian/diff-backup/internal/db_mng"
	"github.com/looCiprian/diff-backup/internal/file_mng"
	"os"
)

type initCommand struct {
	destination string
}

var initCommandConfiguration initCommand

func SetInitConfig(destination string)  {
	initCommandConfiguration.destination = destination
}

//executeInit
// Execute directory initialization
func ExecuteInit() error {
	destination := file_mng.AddSlashIfNotPresent(initCommandConfiguration.destination)

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

	home, _ := os.UserHomeDir()
	configPath := home + config.ConfigurationFile
	if !file_mng.FileExists(configPath){
		file_mng.CreateConfigFile(configPath)
	}

	fmt.Println("Backup correctly initialized ")

	return nil
}


