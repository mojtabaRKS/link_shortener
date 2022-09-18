FROM golang:1.19.1-alpine3.16

RUN mkdir -p /app

WORKDIR /app

RUN apk add --no-cache git curl wget