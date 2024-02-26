# gRPC Consumer / Producer
## Bootstrap by team: Portugese Man O' War
Alexander Snit, Daniel Kogan, Dylan Scott, Gretta Halollari

### Running the code
To compile:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc
protoc --go_out=. --go-grpc_out=. peernode.proto
```

To run mock market server:
```bash
go run market/mock.go
```

To run consumer:
```bash
go run consumer/consumer.go
```

To run producer:
```bash
go run producer/producer.go
```

### Testing

The test provided will run the market server, and then make a request with the producer to the market to add a file hash url, then the consumer will query the market for the producer and the producer for the file. The test will check if the consumer and producer received the expected response.

To run the test:
```bash 
./test.sh
```

### Notes

