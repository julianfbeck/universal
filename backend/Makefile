NOW_RFC3339 = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

all: build cmddocs


runctr:
	@echo "Running control"
	go run ./control/cmd/main.go

runwkr:
	@echo "Running worker"
	go run ./worker/cmd/main.go

package: 
	@echo "Packaging"
	go build -o ./bin/control ./control/cmd/main.go
	go build -o ./bin/worker ./worker/cmd/main.go

docker:
	@echo "Building docker image"
	docker build -t $DOCKER_IMAGE .

test:
	go test ./...  -list=.

cover:
	go test -v -coverprofile cover.out  ./... 
	go tool cover -html=cover.out -o cover.html
	open cover.html

docs-server: 
	@echo Running Docs Generation
	cd control/cmd
	godoc -http=:6060 

docs-worker: 
	@echo Running Docs Generation
	cd worker/cmd
	godoc -http=:6060 