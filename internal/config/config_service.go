package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var configuration Configuration

const ConfigurationFile = ".diff-backup-config.yaml"

type Configuration struct {
	Files struct {
		Blacklistnamefile []string
	}
}

func LoadConfiguration(destination string) {

	config, err := ioutil.ReadFile(destination) // just pass the file name
	if err != nil {
		log.Fatalf("error: %v - Use the init option before performing a backup 'init -s <destination directory>'", err)
	}

	err = yaml.Unmarshal(config, &configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

func GetBlackListedFiles() []string {
	return configuration.Files.Blacklistnamefile
}

func CreateConfigFile(destination string) error {

	var defautlConfig = `files:
    blacklistnamefile:
        - ".DS_Store"`

	f, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err1 := f.WriteString(defautlConfig)

	if err1 != nil {
		return err
	}

	return nil
}
