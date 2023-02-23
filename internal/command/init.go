package command

import (
	"errors"
	"fmt"

	"github.com/looCiprian/diff-backup/internal/config"
	"github.com/looCiprian/diff-backup/internal/file_mng"
)

type initCommand struct {
	destination string
}

var initCommandConfiguration initCommand

func SetInitConfig(destination string) {
	initCommandConfiguration.destination = destination
}

// executeInit
// Execute directory initialization
func ExecuteInit() error {
	destination := file_mng.AddSlashIfNotPresent(initCommandConfiguration.destination)

	if !file_mng.IsEmptyDirectory(destination) {
		fmt.Println("The directory is not empty ")
		for i, file := range file_mng.FilesInDirectory(destination) {
			fmt.Println("File" + string(i) + " Name: " + file)
		}
		return errors.New("")
	}

	if file_mng.CreateNewFileWithContent(destination+"IMPORTANT.txt", "DO NOT DELETE / ADD FILES MANUALLY, IF YOU NEED SOME DATA ONLY COPY OPERATIONS ARE ALLOWED") != nil {
		return errors.New("Cannot write important file ")
	}

	configPath := destination + config.ConfigurationFile
	if !file_mng.FileExists(configPath) {
		config.CreateConfigFile(configPath)
	}

	fmt.Println("Backup correctly initialized ")

	return nil
}
