# Go OpenAPI Manage API

## Overview

This project is a Go-based API for managing OpenAPI specifications. It provides endpoints for creating, reading, updating, and deleting OpenAPI specs stored in a PostgreSQL database.

## Features

- Connects to a PostgreSQL database
- RESTful API endpoints for managing OpenAPI specs
- Modular architecture with handlers, routes, and storage layers

## Prerequisites

- Go 1.16 or later
- PostgreSQL database
- [Optional] Docker for containerized deployment

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ycchuang1999/go-openapi-manage-api.git
   cd go-openapi-manage-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your environment variables:

   - `DB_HOST`: Database host
   - `DB_PORT`: Database port
   - `DB_USER`: Database user
   - `DB_PASSWORD`: Database password
   - `DB_NAME`: Database name
   - `PORT`: Port for the API server (default is 8080)

## Usage

1. Run the application:

   ```bash
   go run cmd/go-openapi-manage-api/main.go
   ```

2. Access the API at `http://localhost:8080`.

## API Endpoints

- `GET /specs`: Retrieve all OpenAPI specs
- `POST /specs`: Create a new OpenAPI spec
- `GET /specs/{id}`: Retrieve a specific OpenAPI spec
- `PUT /specs/{id}`: Update a specific OpenAPI spec
- `DELETE /specs/{id}`: Delete a specific OpenAPI spec

