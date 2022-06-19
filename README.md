# JWT Server with [gorilla/mux](https://github.com/gorilla/mux)

## Getting Started
### Run Server
1. Run containers
    ```shell
    make run
    ```
2. Run tests
    ```shell
    make test
    ```
   
3. Stop containers
   ```shell
   make stop
   ```

## API
### POST /signup
```shell
curl --location --request POST 'localhost:8888/signup' \
--form 'email=test@sample.com' \
--form 'password=test12345' \
--form 'username=user1'
```

### POST /login
```shell
curl --location --request POST 'localhost:8888/login' \
--form 'email=test@sample.com' \
--form 'password=test12345'

```

### POST /v1/me
```shell
curl --location --request GET 'localhost:8888/v1/me' \
--header "Authorization: Bearer $JWT_TOKEN"
```
