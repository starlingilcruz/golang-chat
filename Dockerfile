FROM golang:1.21-alpine AS builder

LABEL MAINTAINER="starlin.gil.cruz@gmail.com"

WORKDIR /go/src/github.com/starlingilcruz/golang-chat

COPY . .

RUN apk add --no-cache git

RUN go build -o golangchat

CMD ["./golangchat"]

EXPOSE 8000