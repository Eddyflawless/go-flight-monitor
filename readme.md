
## About Flight Monitor
<hr>
This is a go project to monitor flight routes and other informations using websockets, go and vanilla javascript. 

## Why this project
Still tinkering with the go language whilst creating various case studies to understand the language more. 

## Goals
- Implement websockets with go and vanilla javascript
- Scrap stats for flight using flight api and goroutine
- Store historic flight data in a database like redis 
- Learn swagger with Go

## Requirements
-   Docker installed
-   You will need Go version 1.11+ installed on your development machine
-   An Api key from https://aviationstack.com/. Create a free account to access key.


## Instructions
In the same or root directory you can run these available commands in the terminal

Run the command `mv .sample-env .env` and insert the apikey generated for you after you created an account on https://aviationstack.com/.

### Using go cli
-   `go run src/main.go` to start the server
-   `go build src/main.go && ./main` to build and start the server

###  Have Docker installed? 
-  `make build` to build image
-  `make start` to run application

## Api calls
`curl https://api.aviationstack.com/v1/flights?access_key=YOUR_ACCESS_KEY`

Optional parameters:
- limit = 100
- offset = 0
- callback = my_callback




## Convert any json to go Struct
I found an easy way to construct a struct from a json response. Visit https://mholt.github.io/json-to-go/ to see the tool.
