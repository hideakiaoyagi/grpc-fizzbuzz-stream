package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "../proto"
)

const (
	port = ":50051"
	fizz = "Kuma"
	buzz = "Mon"
)

type server struct{}

func (s *server) GetFizzbuzzList(req *pb.FizzbuzzRequest, stream pb.FizzbuzzService_GetFizzbuzzListServer) error {
	if req.StartNumber < 1 {
		return fmt.Errorf("'StartNumber' must be a natural number")
	}
	if req.EndNumber < 1 {
		return fmt.Errorf("'EndNumber' must be a natural number")
	}
	if req.StartNumber > req.EndNumber {
		return fmt.Errorf("'EndNumber' must be equal to or greater than 'StartNumber'")
	}

	for i := int(req.StartNumber); i <= int(req.EndNumber); i++ {
		var word string
		switch {
		case i%15 == 0:
			word = fizz + buzz
		case i%3 == 0:
			word = fizz
		case i%5 == 0:
			word = buzz
		default:
			word = strconv.Itoa(i)
		}
		if err := stream.Send(&pb.FizzbuzzResponse{Word: fmt.Sprintf("%v", word)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFizzbuzzServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
