FROM golang:1-alpine AS builder
WORKDIR /app
# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download
# Building a binary
COPY . .
RUN go build -o /bin/SdalGIABot ./cmd/main.go

FROM scratch AS bot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/SdalGIABot SdalGIABot
ENTRYPOINT ["./SdalGIABot"]
