package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "../proto"
)

const (
	server = "localhost"
	port   = ":50051"
)

func main() {
	var (
		start int32 = 1
		end   int32 = 20
	)

	conn, err := grpc.Dial(server+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFizzbuzzServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetFizzbuzzList(ctx, &pb.FizzbuzzRequest{StartNumber: start, EndNumber: end})
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	log.Printf("success: received responses from gRPC server.")

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Println(recv.Word)
	}
}
