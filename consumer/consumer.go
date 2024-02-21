package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/daminals/cse416-init-repo-union-1/peernode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	marketAddr         = flag.String("addr", "localhost:50051", "the address to connect to")
	port         = flag.Int("port", 50052, "The consumer service port")
	fileResponse = &fileInfo{} // this is the file info that the producer will send to me
	endServer    = false       // this is a flag to close the server once the file info is received
	srv          *grpc.Server  // this is the grpc server for the producer to reach out to me
)

type server struct {
	pb.UnimplementedConsumerServiceServer // this is the consumer service
}

type fileInfo struct {
	Link           string // url to download the file
	Token          string // access token to download the file
	PaymentAddress string // payment address to send the payment
}

// implement a function in fileinfo to check if the file info is complete
func (f *fileInfo) isComplete() bool {
	return f.Link != "" && f.Token != "" && f.PaymentAddress != ""
}

// IsAlive is the function that the market will call to check if the consumer is alive
func (s *server) IsAlive(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	log.Printf("Received IsAlive: %v", in)
	return &emptypb.Empty{}, nil
}

// RecieveFileInfo is the function that the producer will call to send the file info
// afterwards, the consumer should close the server and make an http request to the producer
// to download the file
func (s *server) ReceiveFileInfo(ctx context.Context, in *pb.FileLink) (*emptypb.Empty, error) {
	// get the ip address of the producer
	peerCtx, _ := peer.FromContext(ctx)
	log.Printf("Received: %v from producer at %s", in, peerCtx.Addr.String())

	fileResponse.Link = in.GetLink()
	fileResponse.Token = in.GetToken()
	fileResponse.PaymentAddress = in.GetPaymentAddress()

	// Close the server
	endServer = true

	// For now, just return an empty response
	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*marketAddr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		// this context dialer is used to specify the source ip address, so that the producer can reach out to me on the same port
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			dst, err := net.ResolveTCPAddr("tcp", addr)
			// check if destination address is valid
			if err != nil {
					return nil, err
			}
			// create a specified source address
			src := &net.TCPAddr{
					IP:   net.ParseIP("127.0.0.1"), // this is the source ip address, change it to 0.0.0.0 in production
					Port: *port,
			}
			return net.DialTCP("tcp", src, dst)
	}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	marketConnection := pb.NewMarketServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = marketConnection.AddFileRequest(ctx, &pb.FileHash{Hash: "hash"})
	if err != nil {
		log.Fatalf("could not add file request: %v", err)
	}
	log.Printf("Sent: file request to market server at %s", *marketAddr)

	// now i open a grpc connection for the producer to reach out to me
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// explicitly close the connection once file request is made
	conn.Close()
	// now can reuse same port to listen for the producer to connect

	// create an async goroutine to listen for closing the server
	go func() {
		// this will listen in the background for the endServer flag to be true
		// once it is true, it will close the server
		for {
			if endServer {
				srv.Stop()
				break
			}
			// wait 1 second
			time.Sleep(1 * time.Second)
		}
	}()

	// create a new grpc server for the producer to reach out to me
	srv = grpc.NewServer()
	pb.RegisterConsumerServiceServer(srv, &server{})
	log.Printf("Consumer Server listening on port %d...\n", *port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// check if the file info is complete
	if !fileResponse.isComplete() {
		// fatal crash if any of the file info is missing
		log.Fatalf("File info is missing")
	}

	// send an http request to the producer to download the file
	netClient := &http.Client{}
	req, err := http.NewRequest("GET", fileResponse.Link, nil)
	if err != nil {
		log.Fatalf("Error creating http request: %v", err)
	}

	// add the access token in the header
	req.Header.Set("Authorization", fileResponse.Token)

	// send the request
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending http request: %v", err)
	}

	// check if the response is 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error downloading file: %v", resp.Status)
	}
}
