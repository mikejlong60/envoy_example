package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "envoy_example/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
	response := &pb.HelloResponse{Greeting: "Hello " + name}
	return response, nil
}

func main() {
	fmt.Printf("1\n")
	address := "0.0.0.0:5050"
	fmt.Printf("2\n")
	lis, err := net.Listen("tcp", address)
	fmt.Printf("3\n")
	if err != nil {
		log.Fatalf("4 Backend Error %v", err)
	}
	fmt.Printf("5 Server is listening on %v ...\n", address)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	fmt.Printf("6\n")

	s.Serve(lis)
	fmt.Printf("7\n")
}
