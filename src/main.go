package main

import(
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	config "github.com/Eddyflawless/flight-monitor/config"
	controller "github.com/Eddyflawless/flight-monitor/controller"
	"log"
)

var configVar *config.ConfigVar 

func setupRoutes() {

	http.HandleFunc("/", controller.IndexPage)
	http.HandleFunc("/ws/flights", controller.FlightWs)
	http.HandleFunc("/flights", controller.FlightHistory)

	port := fmt.Sprintf(":%s", configVar.Port)
	log.Fatal(http.ListenAndServe(port, nil))
	
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	configVar = config.GetVariables()

	setupRoutes() // setupRoutes

}