package file_mng

import (
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"path/filepath"
)

func FileExists(destination string) bool {
	if _, err := os.Stat(destination); err == nil {
		return true
	}
	return false
}

//copyFile
// Copy file from source to destination
func CopyFile(source string, size int64, destination string) (int64, error) {

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

func CreateNewFileWithContent(path string, content string) error  {
	newFile, err := os.Create(path)

	_, err = newFile.WriteString(content)

	defer newFile.Close()

	if err != nil {
		return err
	}
	return nil

}