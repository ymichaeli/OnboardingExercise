#!/bin/bash
# Build
FROM golang:1.18-bullseye AS build

WORKDIR /app/src
COPY go.mod .
COPY go.sum .

COPY cmd cmd
COPY pkg pkg
RUN go build -o /app/bin/server ./cmd/server

# Run
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=build /app/bin/server server
COPY cmd/config/config.json cmd/config/config.json

EXPOSE 8080

CMD ["/app/server"]