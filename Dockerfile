FROM golang:1.17

WORKDIR /go/src/app

RUN apt-get update

RUN go get -u github.com/cosmtrek/air

EXPOSE 8888

CMD air
