FROM golang:1.15

WORKDIR /go/src/app

RUN go get -u github.com/cosmtrek/air

EXPOSE 8888

CMD ["air"]