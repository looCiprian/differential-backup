package command

import (
	"diff-backup/internal/db_mng"
	"diff-backup/internal/file_mng"
	"diff-backup/internal/time_mng"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/kalafut/imohash"
	"os"
	"path/filepath"
)

type restoreCommand struct {
	destination string
	source      string
	date		string
}

func executeRestore(command restoreCommand) error {
	destination := file_mng.AddSlashIfNotPresent(command.destination)
	source := file_mng.AddSlashIfNotPresent(command.source)
	dateFromRestore := command.date
	databasePath := file_mng.AddSlashIfNotPresent(source) + "index.db"

	if !file_mng.FileExists(databasePath) {
		return errors.New("Backup Directory not initialized. Use init option ")
	}

	db_mng.OpenDB(databasePath)

	// Check if restore date exists
	if _, err := db_mng.IsRestoreDateExist(dateFromRestore); err!=nil{
		return errors.New("No Date " + dateFromRestore + " to restore")
	}

	if err := createRestoreTable(dateFromRestore); err !=nil{
		return err
	}

	dates, err := db_mng.GetAvailableRestoreDates()
	if err != nil{
		return err
	}

	// Get datas from selected to the oldest
	datesToRestore, err := getDateRangeToRestore(dates, dateFromRestore)
	if err != nil {
		return err
	}

	// For each date
	for i:= len(datesToRestore)-1; i>=0; i-- {
		pathSource := source + file_mng.AddSlashIfNotPresent(datesToRestore[i])
		// Iterate the backup date directory
		filepath.Walk(pathSource, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				filePathSource := path
				hash, _ := imohash.SumFile(filePathSource)
				hashString := hex.EncodeToString(hash[:])
				relativeFilePath := filePathSource[len(pathSource):]
				// Check if already restore (only relative file path to maintain only the last file version)
				alreadyRestored, err := db_mng.IsFileAlreadyRestored(relativeFilePath)
				if err != nil {
					return err
				}
				if !alreadyRestored {
					destinationPath := destination + relativeFilePath
					if err != nil {
						return err
					}
					fmt.Println("Restoring: " + info.Name())
					_, err = file_mng.CopyFile(filePathSource, info.Size(), destinationPath)
					if err != nil {
						return err
					}
					// Save restored file to temp restore table
					if _, err:= db_mng.AddRestoredFile(databasePath, info.Name(), relativeFilePath, hashString, "", info.Size()); err!=nil{
						return err
					}
				}else {
					fmt.Println("Skipping: " + info.Name() + " already restored")
				}
			}

			return nil
		})
	}

	if err := db_mng.DropRestoreTable();err!=nil{
		return err
	}

	return nil
}

func createRestoreTable(date string) error {
	if _, err := db_mng.CreateTempTable(); err !=nil{
		return errors.New("No Date " + date + " to restore")
	}
	return nil
}

func getDateRangeToRestore(dates []string, startDate string) ([]string, error) {

	sortedDates, err := time_mng.SortStringDates(dates)
	if err != nil {
		return nil, err
	}

	var dateToRestore []string

	for _,date := range sortedDates{
		if date != startDate{
			dateToRestore = append(dateToRestore, date)
		}
		if date == startDate{
			dateToRestore = append(dateToRestore, date)
			break
		}
	}

	return dateToRestore, nil
}