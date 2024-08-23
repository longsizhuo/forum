1. First
```shell
# go version: 1.23.0
go mod tidy
protoc --go_out=. --go-grpc_out=. proto\forum.proto
```