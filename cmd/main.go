package main

import (
	"log"
	"os"

	"github.com/goocarry/wb-internship/internal/config"
	"github.com/goocarry/wb-internship/internal/service"
)

func main() {
	log.Print("info-a9fb8cec: init config")
	cfg := config.GetConfig()

	log.Println("info-1bf4eefa: init logger")
	logger := log.New(os.Stdout, "api_logger: ", log.LstdFlags)

	app, err := service.NewService(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("info-395c0c9c: app is running")
	app.Run()
}
