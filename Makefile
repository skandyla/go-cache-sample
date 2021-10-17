all: lint test
.DEFAULT_GOAL := all

lint:
	golangci-lint run

test:
	go test -v ./...
	go test -race ./...
