package main

import (
	"context"
	pb "encrypted-election-counting/pkg/distributed" // Import the generated protobuf package
	"encrypted-election-counting/pkg/encryption"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type VoteAggregator struct {
	mu          sync.Mutex
	aggregates  map[string][]byte
	encryptFunc func(vote int) []byte
	controller  pb.ControllerClient

	keyShare    []byte
	shareIndex  int32
	totalShares int32
	publicKey   *rlwe.PublicKey

	pb.UnimplementedWorkerServer
}

func (va *VoteAggregator) SubmitVote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	va.mu.Lock()
	defer va.mu.Unlock()

	candidateId, _ := strconv.Atoi(req.GetCandidateId())
	encryptedVote := va.encryptFunc(candidateId)
	if _, exists := va.aggregates[req.CandidateId]; !exists {
		va.aggregates[req.CandidateId] = encryptedVote
	} else {
		va.aggregates[req.CandidateId] = append(va.aggregates[req.CandidateId], encryptedVote...)
	}

	log.Printf("Received vote for candidate %s", req.CandidateId)
	return &pb.VoteResponse{Status: "Vote received"}, nil
}

func (va *VoteAggregator) SendAggregates(ctx context.Context, req *pb.AggregateRequest) (*pb.AckResponse, error) {
	va.mu.Lock()
	defer va.mu.Unlock()

	log.Printf("Sending aggregates to controller: %v", va.aggregates)

	_, err := va.controller.ReceiveAggregates(ctx, &pb.AggregateRequest{EncryptedAggregates: va.aggregates})
	if err != nil {
		log.Printf("Error sending aggregates to controller: %v", err)
		return &pb.AckResponse{Message: "Failed to send aggregates"}, err
	}

	va.aggregates = make(map[string][]byte)
	return &pb.AckResponse{Message: "Aggregates sent successfully"}, nil
}

func (va *VoteAggregator) sendPeriodicAggregates() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := va.SendAggregates(ctx, &pb.AggregateRequest{})
		if err != nil {
			log.Printf("Error during periodic aggregate send: %v", err)
		} else {
			log.Println("Aggregates sent successfully")
		}
	}
}

func (va *VoteAggregator) ReceiveKeyShare(ctx context.Context, req *pb.KeyShareRequest) (*pb.KeyShareResponse, error) {
	va.mu.Lock()
	defer va.mu.Unlock()

	if req.ShareIndex < 0 || req.ShareIndex >= req.TotalShares {
		return &pb.KeyShareResponse{Success: false}, nil
	}

	va.keyShare = req.KeyShare
	va.shareIndex = req.ShareIndex
	va.totalShares = req.TotalShares

	if len(req.PublicKey) > 0 {
		publicKey := &rlwe.PublicKey{}
		if err := publicKey.UnmarshalBinary(req.PublicKey); err != nil {
			return &pb.KeyShareResponse{Success: false}, err
		}
		va.publicKey = publicKey
	}

	log.Printf("Received key share %d of %d", va.shareIndex, va.totalShares)

	va.encryptFunc = func(vote int) []byte {
		if va.publicKey != nil {
			ciphertext, err := encryption.EncryptVote(va.publicKey, int64(vote))
			if err != nil {
				log.Printf("Encryption error: %v", err)
				return nil
			}
			return ciphertext.MarshalBinary()
		}
		return []byte{byte(vote)}
	}
}

func main() {
	controllerAddr := os.Getenv("CONTROLLER_ADDRESS")
	if controllerAddr == "" {
		log.Fatal("CONTROLLER_ADDRESS environment variable must be set")
	}

	conn, err := grpc.NewClient(controllerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to controller: %v", err)
	}
	defer conn.Close()

	controllerClient := pb.NewControllerClient(conn)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	voteAggregator := &VoteAggregator{
		aggregates: make(map[string][]byte),
		encryptFunc: func(vote int) []byte {
			cyphertext, err := encryption.EncryptVote(controllerPublicKey, int64(vote))
			if err != nil {
				log.Printf("Encryption error: %v", err)
				return nil
			}
			cyphertextBytes, err := cyphertext.MarshalBinary()
			if err != nil {
				log.Printf("Marshal error: %v", err)
				return nil
			}
			return cyphertextBytes
		},
		controller: controllerClient,
	}

	pb.RegisterWorkerServer(grpcServer, voteAggregator)

	go voteAggregator.sendPeriodicAggregates()

	log.Println("Worker server is running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50052: %v", err)
	}
}
