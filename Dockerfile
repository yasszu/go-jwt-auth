FROM golang:1.18

WORKDIR /go/src/app

RUN apt-get update && \
    apt-get -y install postgresql

RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/golang/mock/mockgen@v1.6.0

EXPOSE 8888

CMD while ! pg_isready --host=$POSTGRES_HOST --port=$POSTGRES_PORT; do sleep 1; done && air
