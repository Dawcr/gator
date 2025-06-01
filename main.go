package main

import (
	"log"

	"github.com/dawcr/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cfg.Print()

	if err = cfg.SetUser("daw"); err != nil {
		log.Fatalf("error setting username: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cfg.Print()
}
