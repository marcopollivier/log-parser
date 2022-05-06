package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	LoggingEnable bool `json:"LoggingEnable"`
}

func get() *Configuration {
	configurationFile := configurationFile()

	var config Configuration

	if err := json.Unmarshal(configurationFile, &config); err != nil {
		log.Fatalf("Properties could not be converted %v", err)
	}

	return &config
}

func Init() {
	if config := get(); !config.LoggingEnable {
		log.SetOutput(ioutil.Discard)
	}
}
