package config

import (
	"io/ioutil"
	"log"
)

func configurationFile() []byte {
	configurationFile, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		log.Fatalf("Config file could not be read %v", err)
	}

	return configurationFile
}
