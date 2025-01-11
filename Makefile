.PHONY: build run test clean

build:
	go build -o bin/api cmd/go-openapi-manage-api/main.go

run:
	go run cmd/go-openapi-manage-api/main.go

test:
	go test ./...

clean:
	rm -rf bin/
