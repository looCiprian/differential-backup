package file_mng

import (
	"io/ioutil"
	"log"
	"os"
)

func DirectoryExists(directory string) bool {
	if _, err := os.Stat(directory); !os.IsNotExist(err) {
		return true
	}
	return false
}

func CreateDirectoryIfNotExists(destination string) bool {
	if !DirectoryExists(destination) {
		err := os.Mkdir(destination, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false
}

func DirectoriesInPath(destination string) []os.FileInfo {

	files, err := ioutil.ReadDir(destination)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func IsEmptyDirectory(destination string) bool{
	files, err := ioutil.ReadDir(destination)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0{
		return true
	}
	return false
}