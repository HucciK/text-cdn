package main

import (
	"log"
	"os"
	"text-cdn/config"
	"text-cdn/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logfile, err := os.Open("logfile")
	if err != nil {
		panic(err)
	}

	logger := log.New(logfile, "", 0)

	app := app.New(logger, cfg)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
