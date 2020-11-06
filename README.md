# JWT Server with Echo
## Getting Started
1. Start Postgres
    ```
    $ docker-compose up
    ```

1. Run server
    ```
    $ go run main.go
    ```

## API
### POST /signup
```sh
curl --location --request POST 'localhost:8888/signup' \
--form 'email=test@sample.com' \
--form 'password=12345'
```

### POST /login
```sh
curl --location --request POST 'localhost:8888/login' \
--form 'email=test@sample.com' \
--form 'password=12345'

```

### POST /v1/me
```sh
curl --location --request GET 'localhost:8888/v1/me' \
--header 'Cookie: Authorization={your_jwt_token}'
```
