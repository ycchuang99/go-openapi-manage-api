FROM golang:1.21-alpine

WORKDIR /app

# Install PostgreSQL client for migrations
RUN apk add --no-cache postgresql-client

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o /bin/api cmd/go-openapi-manage-api/main.go

# Expose port
EXPOSE 8080

CMD ["/bin/api"] 