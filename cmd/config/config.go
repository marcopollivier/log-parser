package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	LoggingEnable bool `json:"LoggingEnable"`
}

func configurationFile() []byte {
	configurationFile, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		log.Fatalf("Config file could not be read %v", err)
	}

	return configurationFile
}

func Get() *Configuration {
	configurationFile := configurationFile()

	var config Configuration

	if err := json.Unmarshal(configurationFile, &config); err != nil {
		log.Fatalf("Properties could not be converted %v", err)
	}

	return &config
}
