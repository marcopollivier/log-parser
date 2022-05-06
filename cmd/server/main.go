package main

import (
	"io/ioutil"
	"log"

	"github.com/marcopollivier/log-parser/cmd/config"
	"github.com/marcopollivier/log-parser/internal/stdin"
)

func main() {
	log.Print("Starting Application")

	Init()

	stdin.Init()
}

func Init() {
	if config := config.Get(); !config.LoggingEnable {
		log.SetOutput(ioutil.Discard)
	}
}
