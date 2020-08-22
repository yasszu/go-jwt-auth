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
--form 'email=test@sample.com' \
--form 'password=12345'
```

### POST /login
```sh
curl --location --request POST 'localhost:8088/login' \
--form 'email=test@sample.com' \
--form 'password=12345'

```

### POST /v1/verify
```sh
curl --location --request GET 'localhost:8088/v1/verify' \
--header 'Cookie: Authorization={your_jwt_token}'
```