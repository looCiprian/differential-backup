package command

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/looCiprian/diff-backup/internal/config"
	"github.com/looCiprian/diff-backup/internal/file_mng"
	"github.com/looCiprian/diff-backup/internal/time_mng"
)

type backupCommand struct {
	destination string
	source      string
}

type workers struct {
	maxWorkers     int
	currentWorkers int
	wg             sync.WaitGroup
	mu             sync.Mutex
}

var backupCommandConfiguration backupCommand

func SetBackupConfig(source string, destination string) {
	backupCommandConfiguration.source = source
	backupCommandConfiguration.destination = destination
}

// Execute backup
func ExecuteBackup() error {

	destination := backupCommandConfiguration.destination    // /tmp/backup
	destination = file_mng.AddSlashIfNotPresent(destination) // /tmp/backup/
	config.LoadConfiguration(destination)

	destinationRoot := destination
	source := backupCommandConfiguration.source // /tmp/source
	date := time_mng.CurrentDate()
	datePath := date + "/" // 31-12-2021/

	destination = destination + datePath // /tmp/backup/ + 31-12-2021/

	if file_mng.CreateDirectoryIfNotExists(destination) {
		fmt.Println("Created a new backup dir: " + destination)
	}

	// Create a new workers pool
	worker := workers{
		maxWorkers:     5,
		currentWorkers: 0,
		wg:             sync.WaitGroup{},
		mu:             sync.Mutex{},
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Check if there are available workers
			worker.mu.Lock()
			available := availableWorkers(&worker)
			worker.mu.Unlock()

			if !available {
				for {
					if availableWorkers(&worker) {
						break
					}
				}
			}

			// Add a new worker
			worker.mu.Lock()
			worker.currentWorkers++
			worker.mu.Unlock()
			worker.wg.Add(1)

			go walkCopy(info, path, source, destination, destinationRoot, &worker)
		}

		return nil
	})
	worker.wg.Wait()
	fmt.Println("Backup Done! ")
	return nil
}

func walkCopy(info os.FileInfo, path string, source string, destination string, destinationRoot string, worker *workers) {

	defer worker.wg.Done()

	fullSourcePath := path // /tmp/source/1/1.txt
	fileName := filepath.Base(fullSourcePath)
	relativePath := fullSourcePath[len(source)+1:] // 1/1.txt
	hashString := file_mng.GetFileHash(fullSourcePath)
	fileExists, date := isFileAlreadyBackup(destinationRoot, relativePath, hashString)

	// No error no file present in backup
	if !fileExists && !file_mng.BlackListedFile(fileName) {

		fmt.Println("Coping file: " + info.Name())
		_, err := file_mng.CopyFile(fullSourcePath, info.Size(), destination+relativePath)
		if err != nil {
			fmt.Println("Error copying file: " + info.Name())
		}
	} else if fileExists { // No error but file already exists

		fmt.Println("File " + info.Name() + " already present in " + destinationRoot + date + relativePath)
	}

	// Decrease the number of workers
	worker.mu.Lock()
	worker.currentWorkers--
	worker.mu.Unlock()

}

func isFileAlreadyBackup(backupPath string, relativePath string, hash string) (bool, string) {

	directories := file_mng.DirectoriesInPath(backupPath)

	for _, dir := range directories {
		if dir.IsDir() {
			path := backupPath + dir.Name() + "/" + relativePath
			if file_mng.FileExists(path) {
				if file_mng.GetFileHash(path) == hash {
					return true, file_mng.AddSlashIfNotPresent(dir.Name())
				}
			}
		}
	}

	return false, ""
}

func availableWorkers(w *workers) bool {
	return w.currentWorkers < w.maxWorkers
}
