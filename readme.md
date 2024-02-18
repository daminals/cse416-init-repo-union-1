# gRPC Consumer / Producer

To compile:
```bash
go get -u google.golang.org/grpc
protoc --go_out=. --go-grpc_out=. peernode.proto
```

To run mock market server:
```bash
go run market/mock.go
```

To run producer:
```bash
go run producer/producer.go
```

To run consumer:
```bash
go run consumer/consumer.go
```