package main

import (
	"context"
	pb "encrypted-election-counting/pkg/distributed" // Import your protobuf package
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	// Get the controller address from the environment variable
	controllerAddr := os.Getenv("CONTROLLER_ADDRESS")
	if controllerAddr == "" {
		log.Fatal("CONTROLLER_ADDRESS environment variable not set")
	}

	// Connect to the controller gRPC server
	conn, err := grpc.Dial(controllerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to controller: %v", err)
	}
	defer conn.Close()

	client := pb.NewElectionClient(conn)

	// Example: Send a dummy vote (you can replace this with actual worker logic)
	vote := []byte("encrypted-vote")
	ctx := context.Background()
	_, err = client.SubmitVote(ctx, &pb.VoteRequest{EncryptedVote: vote})
	if err != nil {
		log.Fatalf("Failed to submit vote: %v", err)
	}

	log.Println("Vote submitted successfully")
}
