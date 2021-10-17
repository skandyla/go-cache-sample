all: lint test
.DEFAULT_GOAL := all

lint:
	golangci-lint run

test:
	go test -v ./...
	go test -race ./...

cover:
	go test -coverprofile=/tmp/c.out
	go tool cover -func=/tmp/c.out
