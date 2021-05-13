package command

import (
	"diff-backup/internal/db_mng"
	"diff-backup/internal/file_mng"
	"diff-backup/internal/time_mng"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/kalafut/imohash"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"path/filepath"
)

type backupCommand struct {
	destination string
	source      string
}

//executeBackup
// Execute backup
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

			// No error no file present in backup
			if err == nil && !fileExists {
				_, err := db_mng.AddFile(*connection, info.Name(), relativePath, hashString, date)
				// DB error
				if err != nil {
					db_mng.CloseDB(*connection)
					return errors.New("Error: " + err.Error())
				} else { // OK database updated, start copying
					fmt.Println("Coping file: " + info.Name())
					_, err := copyFile(fullSourcePath, info.Size(), destination+relativePath)
					if err != nil {
						// If copy error, rollback DB entry
						_, err := db_mng.DeleteFile(*connection, info.Name(), relativePath, hashString, date)
						if err != nil {
							return err
						}
						return err
					}
				}
			} else if err == nil && fileExists {		// No error but file already exists
				path, dateBackup, err := db_mng.GetFileInDB(*connection, hashString)
				if err == nil {
					fmt.Println("File " + info.Name() + " already present in " + destination + path + " backUp at: " + dateBackup)
				}
			} else if err != nil { // DB error
				 return errors.New("Error: " + err.Error())
			}
		}
		return nil
	})
	return nil
}

//copyFile
// Copy file from source to destination
func copyFile(source string, size int64, destination string) (int64, error) {

	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	// Create new directoty if does not exist
	dir, _ := filepath.Split(destination)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return 0, err
	}
	// Create new file
	newFile, err := os.Create(destination)
	if err != nil {
		return 0, err
	}
	defer newFile.Close()

	bar := progressbar.DefaultBytes(size, "Copying ")

	bytesCopied, err := io.Copy(io.MultiWriter(newFile, bar), sourceFile)
	if err != nil {
		return 0, err
	}
	return bytesCopied, nil
}
