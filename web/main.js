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

            async fetchInitalFlights(){

                try{

                    let response;
                    response = await fetch("http://localhost:8081/flights");
                    console.log("response: " ,response);

                }catch(e){
                    console.error("Failed to fetch init: " + e.message);
                }

            },
            setupWsListeners(){

                if(this.ws){
                    this.ws.onopen = this.wsOnOpen;
                    this.ws.onmessage = this.wsOnMessage;
                    this.ws.onerror = this.wsOnError;
                }
            },
            openFlighItem(flight){
                console.log("-->", flight);
            },
            wsOnOpen(event) {
                console.log("connected...")
            },
            wsOnMessage(event) {

                console.log("message received ==>: ", event)
                const { data } = event;
                this.flights = JSON.parse(data);
                console.log(this.flights)
            
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