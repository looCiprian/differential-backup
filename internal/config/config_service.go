package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configuration Configuration

type Configuration struct {
	Files struct{
		Blacklistnamefile []string
	}
}

func LoadConfiguration()  {
	config, err := ioutil.ReadFile("/Users/lorenzograzian/go/src/diff-backup/internal/config.yaml") // just pass the file name
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