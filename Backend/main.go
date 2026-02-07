package main

import (
	appConfig "ecommerce/config"
	server "ecommerce/internal/api"

	"log"
)

func main() {
	config, err := appConfig.LoadConfig()
	if err != nil{
		log.Fatalf("Error loading config: %v", err)
	}

    server.StartServer(config)
}
