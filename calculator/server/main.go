package main

import (
	"context"
	pb "grpc3/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serv %v\n", err)
	}
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.GetFirstNumber() + in.GetSecondNumber(),
	}, nil
}

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	k := int64(2)
	n := in.GetFirstNumber()
	for n > 1 {

		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	return nil
}
