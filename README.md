# Overview
Implementation of swapi via command line using [gRPC Version](https://github.com/abrucker235/swapi-grpc).

## Generate/Update gRPC
```
protoc -I {yourDirectory}/swapi-cli swapi/swapi.proto --go_out=plugins=grpc:.
```

## Build
```
go build -o swapi-cli
```

## Planets
### Retrieve By Name
```
./swapi-cli grpc planets -n Tatooine
```
### Retrieve By Id
```
./swapi-cli grpc planets -i 1
```