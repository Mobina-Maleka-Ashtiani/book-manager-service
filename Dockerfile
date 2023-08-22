FROM golang:1.20-alpine

ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN mkdir app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

EXPOSE 8080

ENTRYPOINT go run .
