package main

import (
	"log"

	"github.com/marcopollivier/log-parser/infra/config"
	"github.com/marcopollivier/log-parser/internal/stdin"
)

func main() {
	log.Print("Starting Application")

	config.Init()
	stdin.Init()
}
