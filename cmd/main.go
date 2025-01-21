package main

import (
	"fmt"
	"log"

	"github.com/auliardana/fleetify-test/internal/config"
	"github.com/auliardana/fleetify-test/internal/server"
)

func main() {

	// Initialize the configuration
	viperConfig := config.LoadConfig()
	appPort := viperConfig.GetInt("app.port")
	fmt.Println("App Port: ", appPort)

	app := server.InitializeServer()

	err := app.Run(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
