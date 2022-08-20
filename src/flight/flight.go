package flight

import (
	"io/ioutil"
	"os"
	"time"
	"fmt"
	"net/http"
	"encoding/json"
)

type FlightResponse struct {

	Pagination Pagination `json:"pagination"`
	Data []Flights `json:"data"`
}

type Flights struct {

	Flight_date string  `json:"flight_date"`
	Flight_status string `json:"flight_status"`
	Departure Departure `json:"departure"`
	Arrival Arrival `json:"arrival"`
	Airline Airline `json:"airline"`
	Flight Flight `json:"flight"`
	Aircraft interface{} `json:"aircraft"`
	Live interface{} `json:"live"`

}

type Departure struct {

	Airport string `json:"airport"`
	Timezone string `json:"timezone"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
	Terminal interface{} `json:"terminal"`
	Gate interface{} `json:"gate"`
	Delay interface{} `json:"delay"`
	Scheduled time.Time `json:"scheduled"`
	Estimated time.Time `json:"estimated"`
	Actual interface{} `json:"actual"`
	EstimatedRunway interface{} `json:"estimated_runway"`
	ActualRunway interface{} `json:"actual_runway"`

}

type Arrival struct {

	Airport string `json:"airport"`
	Timezone string `json:"timezone"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
	Terminal interface{} `json:"terminal"`
	Gate interface{} `json:"gate"`
	Baggage string `json:"baggage"`
	Delay interface{} `json:"delay"`
	Scheduled time.Time `json:"scheduled"`
	Estimated time.Time `json:"estimated"`
	Actual interface{} `json:"actual"`
	EstimatedRunway interface{} `json:"estimated_runway"`
	ActualRunway interface{} `json:"actual_runway"`

}

type Airline struct {
	Name string `json:"name"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

type Flight struct {
	Number string `json:"number"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
	Codeshared interface{} `json:"codeshared"`
}

type Pagination struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Count int `json:"count"`
	Total int `json:"total"`

}


func GetFlights(endpoint string, options map[string]string) (FlightResponse , error) {

	api_url := fmt.Sprintf("http://api.aviationstack.com/v1/%s",endpoint)

	req, err := http.NewRequest("GET", api_url, nil)

	if err != nil {
		fmt.Println(err)
		return FlightResponse{}, err
	}

	q := req.URL.Query()
	q.Add("access_key", os.Getenv("AVIATION_STACK_API_KEY"))

	//attach options
	for key, value := range options {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return FlightResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return FlightResponse{}, err
	}


	var response FlightResponse

	err = json.Unmarshal(body, &response)
	
	if err != nil {
		fmt.Println(err)
		return FlightResponse{}, err
	}

	fmt.Println("flight response: ",	response)

	return response, nil



}