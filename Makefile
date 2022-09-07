.PHONY=start
CONTAINER_NAME=flight-monitor
NETWORK=flight-monitor

init: 
	- @docker network create $(NETWORK)

start: init dev
	- @echo "All done..." 

start_web:
	- cd web && python3 -m http.server 9000

start_db:
	- @docker-compose -f docker-compose.yaml  up

stop: tear_down
	- @echo "stopping flight monitor application"

build:
	- @docker build -t $(CONTAINER_NAME) .

dev:
	- @docker run --rm --net $(NETWORK) -p8081:8081 -v "./src:/app" -e AVIATION_STACK_API_KEY=api-key $(NETWORK)

swagger:
	- @docker run --rm --net $(NETWORK) -it --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):go -v $HOME:$HOME -w  $(PWD)/src quay.io/goswagger/swagger

tear_down:
	- @docker network rm $(NETWORK)
	- @echo "Removing flight monitor image from local filesystem"
	- @docker stop  $(docker ps -f $(CONTAINER_NAME))