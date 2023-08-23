package main

import (
	"example-data-service/internal/config"

	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config from env
	_ = config.Read()

	return nil
}
