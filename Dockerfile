FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod dowmload

COPY . .

RUN go build
