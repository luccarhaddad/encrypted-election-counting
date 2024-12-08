package main

import (
	"context"
	pb "encrypted-election-counting/pkg/distributed"
	"encrypted-election-counting/pkg/encryption"
	"encrypted-election-counting/pkg/secretsharing"
	"fmt"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ControllerServer struct {
	pb.UnimplementedControllerServer
	mu sync.Mutex

	secretKey *rlwe.SecretKey
	publicKey *rlwe.PublicKey

	totalWorkers int
	threshold    int

	aggregatedVotes map[string]*rlwe.Ciphertext

	keyShares [][]byte
}

func (s *ControllerServer) ReceiveAggregates(ctx context.Context, req pb.AggregateRequest) (*pb.AckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for candidateId, encryptedAggregate := range req.EncryptedAggregates {
		cyphertext := &rlwe.Ciphertext{}
		err := cyphertext.UnmarshalBinary(encryptedAggregate)

		if err != nil {
			return &pb.AckResponse{Message: "Failed to unmarshal ciphertext"}, err
			continue
		}

		if existingCyphertext, exists := s.aggregatedVotes[candidateId]; exists {
			// TODO: Implement homomorphic addition of ciphertexts
			// This requires using Lattigo's Homo-morphic addition methods

		} else {
			s.aggregatedVotes[candidateId] = cyphertext
		}
	}
	return &pb.AckResponse{Message: "Aggregates received successfuly"}, nil
}

func (s *ControllerServer) DecryptFinalResult(ctx context.Context, req *pb.DecryptRequest) (*pb.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reconstructedKeyBytes, err := secretsharing.CombineKey(s.keyShares)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct secret key: %v", err)
	}

	reconstructedKey := &rlwe.SecretKey{}
	err = reconstructedKey.UnmarshalBinary(reconstructedKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal reconstructed key: %v", err)
	}

	finalResults := make(map[string]int32)
	for candidateId, encryptedVote := range s.aggregatedVotes {
		decryptedVote, err := encryption.DecryptVote(reconstructedKey, encryptedVote)
		if err != nil {
			log.Printf("Error decrypting vote for candidate %s: %v", candidateId, err)
			continue
		}
		finalResults[candidateId] = int32(decryptedVote)
	}
	return &pb.ResultResponse{FinalResults: finalResults}, nil
}
func initializeController(totalWorkers, threshold int) (*ControllerServer, error) {
	secretKey, publicKey, err := encryption.GenerateKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to generate keys: %v", err)
	}

	secretKeyBinary, err := secretKey.MarshalBinary()
	keyShares, err := secretsharing.SplitKey(secretKeyBinary, totalWorkers, threshold)

	if err != nil {
		return nil, fmt.Errorf("failed to split key: %v", err)
	}

	return &ControllerServer{
		secretKey:       secretKey,
		publicKey:       publicKey,
		totalWorkers:    totalWorkers,
		threshold:       threshold,
		aggregatedVotes: make(map[string]*rlwe.Ciphertext),
		keyShares:       keyShares,
	}, nil
}

func main() {
	totalWorkers, err := getEnvInt("TOTAL_WORKERS", 5)
	if err != nil {
		log.Fatalf("Failed to get total workers: %v", err)
	}

	threshold, err := getEnvInt("THRESHOLD", 3)
	if err != nil {
		log.Fatalf("Failed to get threshold: %v", err)
	}

	controllerServer, err := initializeController(totalWorkers, threshold)
	if err != nil {
		log.Fatalf("Failed to initialize controller: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterControllerServer(grpcServer, controllerServer)

	go distributeKeyShares(controllerServer.keyShares)

	log.Println("Controller server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
	}
}

func distributeKeyShares(keyShares [][]byte) {

	workerAddress := strings.Split(os.Getenv("WORKER_ADDRESSES"), ",")

	for i, share := range keyShares {
		if i >= len(workerAddress) {
			log.Printf("Not enough workers to distribute key share %d", i)
			break
		}

		conn, err := grpc.NewClient(workerAddress[i], grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to worker %d at %s: %v", i, workerAddress[i], err)
			continue
		}
		defer conn.Close()

		workerClient := pb.NewWorkerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = workerClient.ReceiveKeyShare(ctx, &pb.KeyShareRequest{
			KeyShare:    share,
			ShareIndex:  int32(i),
			TotalShares: int32(len(keyShares)),
		})
		if err != nil {
			log.Printf("Failed to send key share to worker %d: %v", i, err)
		}
	}

}

func getEnvInt(key string, defaultValue int) (int, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s: %v", key, err)
	}
	return value, nil
}
