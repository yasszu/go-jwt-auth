# JWT Server with Echo
## Getting Started
1. Start Postgres
    ```
    $ docker-compose up
    ```

1. Run server
    ```
    $ go run server.go
    ```

## API
### POST /signup
```sh
curl --location --request POST 'localhost:8088/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@sample.com",
    "password": "abcd123"
}'
```

### POST /login
```sh
curl --location --request POST 'localhost:8088/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@sample.com",
    "password": "abcd123"
}'

```

### POST /verify
```sh
curl --location --request POST 'localhost:8088/curl --location --request POST 'localhost:8088/verify' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2Mjg1MTM1ODYsImlhdCI6MTU5Njk3NzU4Niwic3ViIjoxMSwiZW1haWwiOiJ0ZXN0QHNhbXBsZS5jb20ifQ.wU2iZqoggf5QHYjBXlHVdNI4OybxBYEWLGqJHYsbf2s'
```
