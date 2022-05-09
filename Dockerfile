# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ./ ./

RUN apk add build-base
RUN go mod download
RUN go get github.com/mattn/go-sqlite3
RUN go build -o /go-rest-api

EXPOSE 8080

CMD /go-rest-api