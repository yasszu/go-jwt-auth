# JWT Server with Echo

## Getting Started
### Setup environment
1. Install direnv and [set hook at your shell](https://github.com/direnv/direnv/blob/master/docs/hook.md).
    ```
    $ brew install direnv
    $ echo export JWT_SECRET="{your_cecret_key}" > .envrc
    $ direnv allow .
    ```
    
### Run Server
1. Run containers
    ```
    $ docker-compose up
    ```

## API
### POST /signup
```sh
curl --location --request POST 'localhost:8888/signup' \
--form 'email=test@sample.com' \
--form 'password=test12345' \
--form 'user1'
```

### POST /login
```sh
curl --location --request POST 'localhost:8888/login' \
--form 'email=test@sample.com' \
--form 'password=test12345'

```

### POST /v1/me
```sh
curl --location --request GET 'localhost:8888/v1/me' \
--header 'Authorization: Bearer {your_jwt_token}'
```
