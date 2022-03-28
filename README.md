# POC para testar o Go com Redis

## Infra

- Para subir a infra entrar no dir /infra. Necessário ter docker-compose
```
docker-compose up -d 
```

## Run test


- UNIT
 ``` 
 go test github.com/joleques/go-mongo-poc/test/unit
 ``` 

- Integration

``` 
go test github.com/joleques/go-mongo-poc/test/integration
```