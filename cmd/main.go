package main

import (
	"log"
	"os"

	"github.com/Ivanf1/esp-ble-mesh-sync-service/pkg/api"
	"github.com/Ivanf1/esp-ble-mesh-sync-service/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	server := api.NewServer(os.Getenv("SERVER_LISTEN_ADDRESS"))
	log.Println("server running on port:", os.Getenv("SERVER_LISTEN_ADDRESS"))
	log.Fatal(server.Start())
}
