package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/Eddyflawless/flight-monitor/config"
	controller "github.com/Eddyflawless/flight-monitor/controller"
	redis "github.com/Eddyflawless/flight-monitor/database"
	"github.com/Eddyflawless/flight-monitor/flight"
	"github.com/joho/godotenv"
)

var configVar *config.ConfigVar

func setupRoutes() {

	http.HandleFunc("/", controller.IndexPage)
	http.HandleFunc("/ws/flights", controller.FlightWs)
	http.HandleFunc("/flights", controller.FlightHistory)

	port := fmt.Sprintf(":%s", configVar.Port)
	log.Println(http.ListenAndServe(port, nil))

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	rdb, err := redis.New()

	if err != nil {
		panic(err)
	}

	redis.DB = rdb

	configVar = config.GetVariables()
}

func main() {

	go flight.StartCron()
	setupRoutes() // setupRoutes

}
