package config

import (
	"io/ioutil"
	"log"

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

	configPath := destination + ConfigurationFile
	config, err := ioutil.ReadFile(configPath) // just pass the file name
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
