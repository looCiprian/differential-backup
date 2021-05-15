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
	destination := backupcommand.destination					// /tmp/backup
	destination = file_mng.AddSlashIfNotPresent(destination)	// /tmp/backup/
	databasePath := destination + "index.db"	// /tmp/backup/index.db
	source := backupcommand.source			// /tmp/source
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
			relativePath := fullSourcePath[len(source) - len(baseSourcePath)+1:]  // source/1/1.txt
			hash, _ := imohash.SumFile(fullSourcePath)
			hashString := hex.EncodeToString(hash[:])
			fileExists, err := db_mng.IsFileAlreadyBackup(relativePath, hashString, info.Size())

			// No error no file present in backup
			if err == nil && !fileExists {
				_, err := db_mng.AddFile(databasePath, info.Name(), relativePath, hashString, date, info.Size())
				// DB error
				if err != nil {
					return errors.New("Error: " + err.Error())
				} else { // OK database updated, start copying
					fmt.Println("Coping file: " + info.Name())
					_, err := copyFile(fullSourcePath, info.Size(), destination+relativePath)
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
				path, dateBackup, err := db_mng.GetFileInDB(hashString)
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

//copyFile
// Copy file from source to destination
func copyFile(source string, size int64, destination string) (int64, error) {

	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}


	// Create new directory if does not exist
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


	bar := progressbar.DefaultBytes(size, "Progress")

	bytesCopied, err := io.Copy(io.MultiWriter(newFile, bar), sourceFile)
	sourceFile.Close()
	newFile.Close()
	if err != nil {
		return 0, err
	}
	return bytesCopied, nil
}
