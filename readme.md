# gRPC Consumer / Producer

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

The test provided will run the market server, and then make a request with the consumer, and then the producer will make a request to the market server. The test will check if the consumer and producer received the expected response.

To run the test:
```bash 
./test.sh
```

### Notes

We will likely want to establish a communication format for the market to test if the consumer is still active, and if not, remove it from the list of active consumers. Perhaps there can also be a reporting mechanism for the producer to report if the consumer is not responding to requests.

This can be implemented by the market team. The IsAlive protocol is created for this purpose, and is implemented in the consumer. The producer does not need to call IsAlive since it will be requesting via the ReceiveFileInfo protocol.