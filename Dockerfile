FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY . .

RUN apk add protoc

# RUN go mod download

# RUN go build -o ./cmd/link_shortener/main

# CMD ./cmd/link_shortener/main
