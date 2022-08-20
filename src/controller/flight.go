package controller

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	flight "github.com/Eddyflawless/flight-monitor/flight"
	websocket "github.com/Eddyflawless/flight-monitor/websocket"
)


func FlightHistory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		options := make(map[string]string)

		flightResponse, err := flight.GetFlights("flights" , options)

		if err != nil {
			log.Fatal(err)
		}

		flights, err := json.Marshal(flightResponse)

		if err != nil {
			panic(err)
		}

		fmt.Println(flightResponse)

		w.Header().Set("Content-Type", "application/json")
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

func IndexPage( w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Fprintf(w, "index page2..")

	}else{
		w.WriteHeader(405)
	
		w.Write([]byte("Method not supported"))
	}
}

