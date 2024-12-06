package main

import (
	"context"
	pb "encrypted-election-counting/pkg/distributed"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedElectionServer
}

func (s *server) SubmitVote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received encrypted vote: %v", req.EncryptedVote)
	return &pb.VoteResponse{Status: "Vote received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterElectionServer(grpcServer, &server{})

	log.Println("Starting controller server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
