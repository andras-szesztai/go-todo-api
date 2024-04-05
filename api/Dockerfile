# syntax=docker/dockerfile:1


FROM golang:1.22.1 AS build-stage

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./api/

WORKDIR /app/api
RUN go mod download

COPY . .

RUN pwd

CMD ["air", "-c", ".air.toml"]