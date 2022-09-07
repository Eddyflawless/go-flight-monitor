package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	flight "github.com/Eddyflawless/flight-monitor/flight"
	"github.com/gorilla/websocket"
)

var ctx = context.Background()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}

	return conn, err
}

func Writer(conn *websocket.Conn) {

	for {

		ticker := time.NewTicker(time.Second * 10) // every 10 seconds

		for t := range ticker.C {

			fmt.Println("updating subscribed views", t)
			// options := make(map[string]string)

			// flightResponse, err := flight.GetFlights("flights" , options)
			flightResponse, err := flight.GetFlights(ctx)

			if err != nil {
				log.Fatal(err)
			}

			flights, err := json.Marshal(flightResponse)

			if err != nil {
				panic(err)
			}

			//save to database
			fmt.Println(ctx)
			// flight.SaveFlights(ctx,flights)

			if err = conn.WriteMessage(websocket.TextMessage, []byte(flights)); err != nil {
				fmt.Println("error writing message", err)
				return
			}

		}
	}

}
