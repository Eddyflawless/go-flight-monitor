package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var ctx = context.Background()

func scrapeThirdPartyEndpoint(endpoint string, options map[string]string) (FlightResponse, error) {

	api_url := fmt.Sprintf("http://api.aviationstack.com/v1/%s", endpoint)

	// api_url := fmt.Sprintf("%s/%s", config.GetVariables().ApiUrlDev, endpoint)

	fmt.Printf("endpoint: %s\n", endpoint)
	fmt.Printf("accessing %v \n", api_url)

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

	fmt.Println("flight response: ", response)

	return response, nil

}

func StartCron() {

	for {

		ticker := time.NewTicker(time.Second * 300) // every x seconds

		for t := range ticker.C {
			options := make(map[string]string)

			// scrap flight and insert into database
			flightResponse, err := scrapeThirdPartyEndpoint("flights", options)

			if err != nil {
				log.Fatal(err)
			}

			log.Print(t)

			// log.Print("flight data: ", flights)

			log.Print("flight data: ", flightResponse)

			log.Println("insert into redis database")
			SaveFlights(ctx, flightResponse.Data)

		}
	}

}
