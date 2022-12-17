"use strict";

function Init(){
    bootstrapVue();
    // bootstrapWs();
}


function bootstrapVue() {

    const { createApp } = Vue;

    createApp({
        data() {
            return {
                ws: null,
                flights: [],
            }
        },
        methods: {


            poplulateFlights(data){
                if(!data) return

                this.flights = data.map((flight, i) =>  { flight.isShow = false; flight.id = i+1; return flight} );

                console.log(this.flights)

            },
            async fetchInitalFlights(){

                try{

                    let data;
                    data = await fetch("http://localhost:8081/flights").then((response) => response.json());
                    if(!data) return

                    this. poplulateFlights(data)

                }catch(e){
                    console.error("Failed to fetch init: " + e.message);
                }

            },
            notAvailableData(field){
                return field? field: "N/A";
            },
            setupWsListeners(){

                if(this.ws){
                    this.ws.onopen = this.wsOnOpen;
                    this.ws.onmessage = this.wsOnMessage;
                    this.ws.onerror = this.wsOnError;
                }
            },
            openFlighItem(flight){
                flight.isShow = !flight.isShow ;
            },
            wsOnOpen(event) {
                console.log("connected...")
            },
            wsOnMessage(event) {

                let data ;
                try{

                    console.log("message received ==>: ", event)
                    const { data } = event;
                    var parsedData = JSON.parse(data);                    
                    this.poplulateFlights(parsedData)


                }catch(e){
                    console.log("error", e);
                }finally{
                    if(data) this.flights = data
                }
            
            },
            wsOnerror(event) {
                console.log(event)
            }
        },
        created(){
            console.log("created...")
            this.fetchInitalFlights();
        },
        mounted() {
            this.ws = new WebSocket("ws://localhost:8081/ws/flights");
            this.setupWsListeners();
        }
    }).mount("#app")
}


Init();