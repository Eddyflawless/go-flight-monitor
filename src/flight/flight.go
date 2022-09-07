package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis "github.com/Eddyflawless/flight-monitor/database"
	redisO "github.com/go-redis/redis"
)

type FlightResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Flights  `json:"data"`
}

type Flights struct {
	Flight_date   string      `json:"flight_date"`
	Flight_status string      `json:"flight_status"`
	Departure     Departure   `json:"departure"`
	Arrival       Arrival     `json:"arrival"`
	Airline       Airline     `json:"airline"`
	Flight        Flight      `json:"flight"`
	Aircraft      interface{} `json:"aircraft"`
	Live          interface{} `json:"live"`
}

type Departure struct {
	Airport         string      `json:"airport"`
	Timezone        string      `json:"timezone"`
	Iata            string      `json:"iata"`
	Icao            string      `json:"icao"`
	Terminal        interface{} `json:"terminal"`
	Gate            interface{} `json:"gate"`
	Delay           interface{} `json:"delay"`
	Scheduled       time.Time   `json:"scheduled"`
	Estimated       time.Time   `json:"estimated"`
	Actual          interface{} `json:"actual"`
	EstimatedRunway interface{} `json:"estimated_runway"`
	ActualRunway    interface{} `json:"actual_runway"`
}

type Arrival struct {
	Airport         string      `json:"airport"`
	Timezone        string      `json:"timezone"`
	Iata            string      `json:"iata"`
	Icao            string      `json:"icao"`
	Terminal        interface{} `json:"terminal"`
	Gate            interface{} `json:"gate"`
	Baggage         string      `json:"baggage"`
	Delay           interface{} `json:"delay"`
	Scheduled       time.Time   `json:"scheduled"`
	Estimated       time.Time   `json:"estimated"`
	Actual          interface{} `json:"actual"`
	EstimatedRunway interface{} `json:"estimated_runway"`
	ActualRunway    interface{} `json:"actual_runway"`
}

type Airline struct {
	Name string `json:"name"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

type Flight struct {
	Number     string      `json:"number"`
	Iata       string      `json:"iata"`
	Icao       string      `json:"icao"`
	Codeshared interface{} `json:"codeshared"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

func GetFlights(ctx context.Context) ([]Flights, error) {

	var flights []Flights

	keys, err := redis.DB.Keys("*").Result()

	if err != nil {
		log.Fatal(err)
		return []Flights{}, err
	}

	for _, key := range keys {

		flight, err := GetFlight(ctx, key)

		if err != nil {
			return []Flights{}, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func GetFlight(ctx context.Context, id string) (Flights, error) {

	var flight Flights

	value, err := redis.DB.Get(id).Result()

	if err != nil {
		log.Fatal(err)
		return Flights{}, err
	}

	if err != redisO.Nil {
		err = json.Unmarshal([]byte(value), &flight)
	}

	return flight, err

}

func createKey(flight Flights) string {

	// airline.departure.desination

	return fmt.Sprintf("%s.%s.%s", flight.Airline.Name, flight.Departure.Airport, flight.Arrival.Airport)

}

func SaveFlight(ctx context.Context, flight Flights) {

	flightByte, err := json.Marshal(flight) // validate flight structure

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", flightByte)

	redis_key := createKey(flight)

	err = redis.DB.Set(redis_key, flightByte, 0).Err()

	if err != nil {
		panic(err)
	}

}

func SaveFlights(ctx context.Context, flights []Flights) {

	for _, flight := range flights {

		SaveFlight(ctx, flight)

	}

	log.Printf("Records %v inserted into the database", len(flights))

}
