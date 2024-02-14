# grpc server

to run
```bash
go get -u google.golang.org/grpc
protoc --go-grpc_out=. example.proto
go run main.go
```