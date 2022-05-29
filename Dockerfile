FROM golang:1.18

WORKDIR /go/src/app

RUN apt-get update

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8888

CMD air
