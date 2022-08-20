package main

import(
	"fmt"
	"net/http"
	// "github.com/Eddyflawless/flight-monitor/flight"
	"log"
)

func appHeaders( w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
}

func flightHistory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		appHeaders(w)
		fmt.Fprintf(w, "")
	}

}

func flightApi(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		appHeaders(w)
		fmt.Fprintf(w, "")
	}

}

func IndexPage( w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Fprintf(w, "index page2..")

	}else{
		w.WriteHeader(405)
	
		w.Write([]byte("Method not supported"))
	}
}

func setupRoutes() {

	http.HandleFunc("/", IndexPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func main() {
	setupRoutes() // setupRoutes
}