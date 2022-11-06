package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "auth-ms/auth.utils/auth"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "port to run the server on")
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Authenticate(ctx context.Context, req *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	log.Printf("Received: %v", req.GetToken())
	return &pb.AuthenticationResponse{UserId: 1, TokenValid: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalln("Unable to start the server on port", *port)
	}

	grpc_server := grpc.NewServer()
	pb.RegisterAuthServer(grpc_server, &server{})
	log.Println("Started server at", lis.Addr())

	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalln("Failed to start server at", lis.Addr())
	}
}
