# gingorm1

## MVC
### presentation -> business -> model

## clean
### model -> presentation -> business

## initial setup

```
mysql -u root -p -h 127.0.0.01 -P 3306
CREATE DATABASE gingorm1;

docker-compose up

air
```

## example manual test

```
curl -d '{"email":"test@gmail.com","password":"test123"}' -X POST http://localhost:3000/guest/register
```

## TODO

- censor logs
- add metrics
- golangci lint
- godotenv to load config from env
