"use strict";

function Init(){
    bootstrapVue();
    bootstrapWs();
}


function bootstrapVue() {

    const { createApp } = Vue;

    createApp({
        data() {
            return {
                message: 'Hello world'
            }
        }
    }).mount("#app")
}

function bootstrapWs() {

    console.log("here...")

    let flights = [];

    const websocket = new WebSocket("ws://localhost:8081/ws/flights");

    websocket.onopen = function(event) {

        console.log("connected...")

    }

    websocket.onerror = function(err) {
        console.log(err)

    }

    websocket.onmessage = function(event)  {
        console.log("message received: ", event)

    }

}

Init();