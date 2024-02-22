package internal

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/grpc/peer"
	// "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedConsumerServiceServer // this is the consumer service
}

var (
	CurrentFileLink *pb.FileLink = &pb.FileLink{}
	serverConsumer  *grpc.Server = nil
)

func SendFileRequest(marketServerAddr string) error {
	// Set up a connection to the server.
	log.Printf("Connecting to market server at %s...", marketServerAddr)
	connMarketServer, err := grpc.Dial(marketServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		// this context dialer is used to specify the source ip address, so that the producer can reach out to me on the same port
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			// check if destination address is valid
			dst, err := net.ResolveTCPAddr("tcp", addr)
			if err != nil {
				return nil, err
			}

			// create a specified source address
			src := &net.TCPAddr{
				IP:   net.ParseIP(ConsumerAddr),
				Port: int(ConsumerPort),
			}
			return net.DialTCP("tcp", src, dst)
		}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	clientMarketServer := pb.NewMarketServiceClient(connMarketServer)
	defer connMarketServer.Close() // close connection after function ends

	// Contact the server and print out its response.
	ctxMarketServer, cancelMarketServer := context.WithTimeout(context.Background(), time.Second)
	defer cancelMarketServer() // cancel the context after function ends
	_, err = clientMarketServer.AddFileRequest(ctxMarketServer, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not add file request: %v", err)
	}
	log.Printf("Sent: file request to market server at %s", marketServerAddr)
	return nil
}

// func StartListener() {
// 	// now i open a port connection for the producer to reach out to me via grpc
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", ConsumerPort))
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}
// 	defer listener.Close()

// 	serverConsumer = grpc.NewServer()
// 	pb.RegisterConsumerServiceServer(serverConsumer, &server{})
// 	log.Printf("Consumer Server listening on port %s...\n", listener.Addr().String())
// 	if err := serverConsumer.Serve(listener); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }

// func StopListener() {
// 	serverConsumer.GracefulStop()
// }
