FROM golang:alpine

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /usr/src/data-integration

COPY go.mod go.sum ./

RUN go mod download

COPY . .