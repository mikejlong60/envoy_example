package main

import (
	"context"
	"fmt"
	"log"

	pb "envoy_example/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:1337", grpc.WithInsecure())
	fmt.Printf("1\n")
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewHelloServiceClient(cc)
	fmt.Printf("2\n")
	request := &pb.HelloRequest{Name: "brian"}
	fmt.Printf("3\n")

	ctx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Bearer foo", "Bar", "baz", "user_dn", "cn=fred,ou=simon")
	fmt.Printf("4\n")

	resp, err := client.Hello(ctx, request)
	fmt.Printf("5\n")
	if err != nil {
		errStatus, isGrpcErr := status.FromError(err)
		if !isGrpcErr {
			fmt.Printf("Unknown error! %v", errStatus.Message())
			return
		}
		code := errStatus.Code()
		msg := errStatus.Message()
		fmt.Println(code)
		fmt.Println(msg)
	} else {
		fmt.Printf("6\n")
		fmt.Printf("Receive response => [%v]", resp.Greeting)
	}
}
