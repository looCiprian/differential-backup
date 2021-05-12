package command

import (
	"diff-backup/internal/db_mng"
	"diff-backup/internal/file_mng"
	"diff-backup/internal/time_mng"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/kalafut/imohash"
	"io"
	"log"
	"os"
	"path/filepath"
)

type backupCommand struct {
	destination string
	source      string
}

func executeBackup(backupcommand backupCommand) error {
	destination := backupcommand.destination
	destination = file_mng.AddSlashIfNotPresent(destination)
	databasePath := destination + "index.db"
	source := backupcommand.source
	date := time_mng.CurrentDate()
	datePath := date + "/"
	destination = destination + datePath

	if !file_mng.FileExists(databasePath) {
		return errors.New("Backup Directory not initialized. Use init option ")
	}
	connection, err := db_mng.OpenDB(databasePath)
	if err != nil {
		return errors.New("Error opening DB ")
	}

	if file_mng.CreateDirectoryIfNotExists(destination) {
		fmt.Println("Created a new backup dir: " + destination)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fullSourcePath := path
			relativePath := path[len(source)+1:]
			hash, _ := imohash.SumFile(fullSourcePath)
			hashString := hex.EncodeToString(hash[:])
			fileExists, err := db_mng.IsFileAlreadyBackup(*connection, relativePath, hashString)

			if err == nil && !fileExists {
				_, err := db_mng.AddFile(*connection, info.Name(), relativePath, hashString, date)
				if err != nil {
					fmt.Println("Error: " + err.Error())
					db_mng.CloseDB(*connection)
				} else {
					fmt.Println("Coping file: " + info.Name())
					copyFile(fullSourcePath, destination+relativePath)
				}
			} else if err == nil && fileExists {
				path, dateBackup, err := db_mng.GetFileInDB(*connection, hashString)
				if err == nil {
					fmt.Println("File " + info.Name() + " already present in " + destination + path + " backUp at: " + dateBackup)
				}
			} else if err != nil {
				fmt.Println("Error: " + err.Error())
			}
		}
		return nil
	})
	return nil
}

func copyFile(source string, destination string) (int64, error) {

	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	defer sourceFile.Close()

	// Create new directoty if does not exist
	dir, _ := filepath.Split(destination)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	// Create new file
	newFile, err := os.Create(destination)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	return bytesCopied, nil
}

func ExecuteBackupTest() {
	//backup --source /Users/lorenzograzian/go/src/diff-backup/test/toBackup --destination /Users/lorenzograzian/go/src/diff-backup/test/backup
	executeBackup(backupCommand{destination: "/Users/lorenzograzian/go/src/diff-backup/test/backup", source: "/Users/lorenzograzian/go/src/diff-backup/test/toBackup"})
}
