FROM golang:1.16

WORKDIR /go/src/app

RUN apt-get update && \
    apt-get -y install postgresql

RUN go get -u github.com/cosmtrek/air && \
    go get -v github.com/rubenv/sql-migrate/...

EXPOSE 8888

CMD while ! pg_isready --host=$POSTGRES_HOST --port=$POSTGRES_PORT; do sleep 1; done && air