package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var configuration Configuration

const ConfigurationFile = "/.diff-backup-config.yaml"

type Configuration struct {
	Files struct{
		Blacklistnamefile []string
	}
}

func LoadConfiguration()  {

	home, err := os.UserHomeDir()
	configPath := home + ConfigurationFile
	config, err := ioutil.ReadFile(configPath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	err = yaml.Unmarshal(config, &configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

func GetBlackListedFiles() []string{
	return configuration.Files.Blacklistnamefile
}