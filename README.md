# gingorm1

This is example of [how to structure your golang project](https://kokizzu.blogspot.com/2022/05/how-to-structure-layer-your-golang-project.html) article with gin and gorm (you can change it to whatever framework and persistence libraries you like, the structure should still be similar). echo version [here](https://github.com/kokizzu/echogorm1/). Personally I won't use function injection for database (managed dependency), I would only use function injection for 3rd party (unmanaged dependency), so this repo only show how to create function injection just for sake of example. To test I would prefer using [dockertest](//github.com/kokizzu/dockertest) so without function injection for database like in [fiber1](//github.com/kokizzu/fiber1), this would violate the clean architecture approach because we embed the database provider directly instead of injecting the dependency, but that's the simplest aproach especially if you are working alone.

```
# MVC
presentation -calls-> business -calls-> model

# Clean
model -injected-into-> business
presentation -calls-> business

presentation should only care about transport and serialization/deserialization
model should only care about DAO and persistence (can be decoupled)
business should only care about business logic use cases

presentation can access business
business can access model

model should not ever depend on business
business should not ever depend on presentation
```

## initial setup

```
mysql -u root -p -h 127.0.0.01 -P 3306
CREATE DATABASE gingorm1;

docker-compose up

air

make test
make testv
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
