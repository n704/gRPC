# GO gRPC Example Square
This example is based on command line tool

# Installation
Installing all dependency files are managed by `dep` package manager
```shell
go get -u github.com/golang/dep/cmd/dep
dep ensure
```
# Working
Example I have written uses a server client communication using `tcp`
Client invoke server's service to get result.
In this example Server does 2 function 
* Square
  - Return square of number
  - default number is 1
* Cube
  - Return cube of number
  - default number is 1
  
```shell
go run client/main.go  -square 122 -cube 9
2018/04/26 15:10:54 Square of 122: 14884
2018/04/26 15:10:54 Cube of 9: 729
```


# Runing
For this to work we need to run first run `server` and `client`.

## Server
```shell
go run server/main.go
```

## Client
```shell
go run client/main.go
```
