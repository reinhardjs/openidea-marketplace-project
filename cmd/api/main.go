package main

import (
	"fmt"

	"github.com/openidea-marketplace/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
		Log: log,
	}, viperConfig)

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}
