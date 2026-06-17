FROM golang:1-alpine

WORKDIR /app

# Installing dev tool: air (hot-reload)
RUN go install github.com/air-verse/air@latest

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying sources
COPY . .

# Run with hot-reload
ENTRYPOINT ["air"]
