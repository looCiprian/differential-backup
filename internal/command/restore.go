package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/looCiprian/diff-backup/internal/config"
	"github.com/looCiprian/diff-backup/internal/file_mng"
	"github.com/looCiprian/diff-backup/internal/time_mng"
)

type restoreCommand struct {
	destination string
	source      string
	date        string
}

var restoreCommandConfiguration restoreCommand

func SetRestoreConfig(source string, destination string, date string) {
	restoreCommandConfiguration.source = source
	restoreCommandConfiguration.destination = destination
	restoreCommandConfiguration.date = date
}

func ExecuteRestore() error {

	destination := file_mng.AddSlashIfNotPresent(restoreCommandConfiguration.destination) // /destination/
	source := file_mng.AddSlashIfNotPresent(restoreCommandConfiguration.source)           // /source/
	config.LoadConfiguration(source)
	dateFromRestore := restoreCommandConfiguration.date // YYYY-MM-DD

	if !file_mng.IsEmptyDirectory(destination) {
		return errors.New("Restore directory (" + destination + ") not empty ")
	}

	// Check if restore date exists
	if !isRestoreDateExist(source + dateFromRestore) {
		return errors.New("No Date " + dateFromRestore + " to restore")
	}

	dates := GetResorableDates()

	// Get datas from the oldest date to the restore date
	datesToRestore, err := getDateRangeToRestore(dates, dateFromRestore)
	if err != nil {
		return err
	}

	// For each date starting from the newest to the oldest
	for i := len(datesToRestore) - 1; i >= 0; i-- {
		pathSource := source + file_mng.AddSlashIfNotPresent(datesToRestore[i])
		// Iterate the backup date directory
		filepath.Walk(pathSource, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				filePathSource := path
				relativeFilePath := filePathSource[len(pathSource):]
				// Check if already restore (only relative file path to maintain only the last file version)
				destinationPath := destination + relativeFilePath
				alreadyRestored := isFileAlreadyRestored(destinationPath)

				if !alreadyRestored {
					destinationPath := destination + relativeFilePath
					if err != nil {
						return err
					}
					_, err = file_mng.CopyFile(filePathSource, info.Size(), destinationPath, "Restoring: "+info.Name())
					if err != nil {
						return err
					}
				} else {
					fmt.Println("Skipping: " + info.Name() + " already restored")
				}
			}

			return nil
		})
	}
	return nil
}

func getDateRangeToRestore(dates []string, startDate string) ([]string, error) {

	sortedDates, err := time_mng.SortStringDates(dates)
	if err != nil {
		return nil, err
	}

	var dateToRestore []string

	for _, date := range sortedDates {
		if date != startDate {
			dateToRestore = append(dateToRestore, date)
		}
		if date == startDate {
			dateToRestore = append(dateToRestore, date)
			break
		}
	}

	return dateToRestore, nil
}

func GetResorableDates() []string {

	source := restoreCommandConfiguration.source

	files := file_mng.DirectoriesInPath(source)

	var dates []string

	for _, file := range files {
		if file.IsDir() {
			dates = append(dates, file.Name())
		}
	}

	return dates
}

func isRestoreDateExist(date string) bool {
	return file_mng.FileExists(date)
}

func isFileAlreadyRestored(filePath string) bool {
	return file_mng.FileExists(filePath)
}
