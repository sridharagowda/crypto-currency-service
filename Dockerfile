FROM golang:1.12.7 as builder

COPY . crypto-currency-service
WORKDIR go/src/crypto-currency-service
ENV GO111MODULE=on

RUN go mod download

COPY . .

EXPOSE 8080