package main

import (
	"log"
	"os"

	"github.com/goocarry/wb-internship/internal/config"
	"github.com/goocarry/wb-internship/internal/service"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := log.New(os.Stdout, "api", log.LstdFlags)

	app, err := service.NewService(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
