package command

import (
	"github.com/looCiprian/diff-backup/internal/config"
	"github.com/looCiprian/diff-backup/internal/db_mng"
	"github.com/looCiprian/diff-backup/internal/file_mng"
	"github.com/looCiprian/diff-backup/internal/time_mng"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/kalafut/imohash"
	"os"
	"path/filepath"
)

type backupCommand struct {
	destination string
	source      string
}

var backupCommandConfiguration backupCommand

func SetBackupConfig(source string, destination string)  {
	backupCommandConfiguration.source = source
	backupCommandConfiguration.destination = destination
}

//executeBackup
// Execute backup
func ExecuteBackup() error {

	config.LoadConfiguration()

	destination := backupCommandConfiguration.destination                  // /tmp/backup
	destination = file_mng.AddSlashIfNotPresent(destination)               // /tmp/backup/
	databasePath := destination + "index.db"                               // /tmp/backup/index.db
	source := backupCommandConfiguration.source                            // /tmp/source
	baseSourcePath := file_mng.AddSlashIfNotPresent(filepath.Base(source)) // source/
	date := time_mng.CurrentDate()
	datePath := date + "/" // 31-12-2021/

	destination = destination + datePath // /tmp/backup/ + 31-12-2021/

	if !file_mng.FileExists(databasePath) {
		return errors.New("Backup Directory not initialized. Use init option ")
	}
	err := db_mng.OpenDB(databasePath)
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
			fullSourcePath := path				// /tmp/source/1/1.txt
			fileName := filepath.Base(fullSourcePath)
			relativePath := fullSourcePath[len(source) - len(baseSourcePath)+1:]  // source/1/1.txt
			hash, _ := imohash.SumFile(fullSourcePath)
			hashString := hex.EncodeToString(hash[:])
			fileExists, err := db_mng.IsFileAlreadyBackup(relativePath, hashString, info.Size())

			// No error no file present in backup
			if err == nil && !fileExists && ! file_mng.BlackListedFile(fileName){
				_, err := db_mng.AddFile(databasePath, info.Name(), relativePath, hashString, date, info.Size())
				// DB error
				if err != nil {
					return errors.New("Error: " + err.Error())
				} else { // OK database updated, start copying
					fmt.Println("Coping file: " + info.Name())
					_, err := file_mng.CopyFile(fullSourcePath, info.Size(), destination+relativePath)
					if err != nil {
						// If copy error, rollback DB entry
						_, err := db_mng.DeleteFile(info.Name(), relativePath, hashString, date)
						if err != nil {
							return err
						}
						return err
					}
				}
			} else if err == nil && fileExists {		// No error but file already exists
				path, dateBackup, err := db_mng.GetFileInDB(relativePath, hashString)
				if err == nil {
					fmt.Println("File " + info.Name() + " already present in " + destination + path + " backUp at: " + dateBackup)
				}
			} else if err != nil { // DB error
				 return errors.New("Error: " + err.Error())
			}
		}
		return nil
	})
	db_mng.CloseDB()
	fmt.Println("Backup Done! ")
	return nil
}