package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	redis "github.com/Eddyflawless/flight-monitor/database"
	flight "github.com/Eddyflawless/flight-monitor/flight"
	websocket "github.com/Eddyflawless/flight-monitor/websocket"
)

var ctx = context.Background()

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func FlightHistory(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		// options := make(map[string]string)
		// flightResponse, err := flight.GetFlights("flights", options)

		id, ok := r.URL.Query()["id"]

		if ok {

			flight_id := id[0]
			flightItem, err := flight.GetFlight(ctx, flight_id)

			if err != nil {

				log.Fatal(err)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("{}"))
			}

			flightBytes, err := json.Marshal(flightItem)

			if err != nil {
				log.Fatal(err)
				w.Write([]byte("{}"))
			}

			w.Write(flightBytes)
			return

		}

		flightResponse, err := flight.GetFlights(ctx)

		if err != nil {
			log.Fatal(err)
		}

		flights, err := json.Marshal(flightResponse)

		if err != nil {
			panic(err)
		}

		fmt.Println(flightResponse)

		w.Write(flights)
	}

}

func FlightWs(w http.ResponseWriter, r *http.Request) {

	fmt.Println("here..")
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)

}

func IndexPage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("controller-->", redis.DB)
	w.Write([]byte("controller"))

	// options := make(map[string]string)

	// flightResponse, err := flight.GetFlights("flights" , options)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// flights, err := json.Marshal(flightResponse)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(flightResponse)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(flights)

}
