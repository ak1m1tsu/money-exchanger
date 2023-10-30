package main

import (
	"log"
	"os"

	"github.com/romankravchuk/currency-exchanger/internal/app"
	"github.com/romankravchuk/currency-exchanger/internal/config"
)

func main() {
	cfg, err := config.New(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
