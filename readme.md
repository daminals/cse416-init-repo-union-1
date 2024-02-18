# grpc server

to run
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc
protoc --go_out=. --go-grpc_out=. example.proto
go run main.go
```