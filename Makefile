GIT_HASH = $(shell git describe --tags --dirty --always)

.PHONY: all run lint test build clean gen-cert

all: clean lint test build run

run: 
	./bin/server

run-tls: gen-cert 
	./bin/server
	
test: lint unit

unit: 
	go test ./... -coverprofile cover.out
	
lint:
	go vet ./...

build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOSUMDB=off go build -o bin/server main.go

docker: test build
	docker build . -f docker/Dockerfile

gen-cert:
	openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out cert.pem -keyout key.pem
	
clean:
	rm -rf cert.pem key.pem bin cover.out