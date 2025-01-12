.PHONY: build run test clean db-create db-migrate db-rollback

build:
	go build -o bin/api cmd/go-openapi-manage-api/main.go

run:
	go run cmd/go-openapi-manage-api/main.go

test:
	go test ./...

clean:
	rm -rf bin/

db-create:
	createdb openapi_manager || true

db-migrate:
	psql -d openapi_manager -f migrations/000001_create_specs_table.up.sql

db-rollback:
	psql -d openapi_manager -f migrations/000001_create_specs_table.down.sql
