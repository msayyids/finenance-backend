package main

import (
	"finenance-app/internal/config"
	"fmt"
	"log"
)

func main() {

	viperConfig := config.NewViper()
	logger := config.Logger()
	db := config.ConnectionDatabase(viperConfig, logger)
	app := config.InitFiber()
	validate := config.NewValidator()
	reddis := config.NewReddisClient()

	config.NewAppConfig(&config.AppConfig{
		App:      app,
		DB:       db,
		Config:   viperConfig,
		Log:      logger,
		Validate: validate,
		Redis:    reddis,
	})

	err := app.Listen(fmt.Sprintf(":%s", viperConfig.GetString("APP_PORT")))
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
