package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func gracefulShutdown() {
	var stop = make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	go func() {
		log.Printf("Caught OS Signal: %+v", <-stop)
		log.Print("Waiting for 2 seconds to finish the application gracefully")

		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}

func Init() {
	if config := get(); !config.LoggingEnable {
		log.SetOutput(ioutil.Discard)
	}

	gracefulShutdown()
}
