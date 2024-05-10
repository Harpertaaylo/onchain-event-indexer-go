// Lightweight blockchain event indexer with gRPC

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type Event struct {
	BlockNumber uint64
	TxHash      string
	Data        string
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	fmt.Println("Connected to Ethereum")

	// gRPC server setup
	grpcServer := grpc.NewServer()
	_ = grpcServer // register services here

	// Simple loop to fetch latest block and print number
	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Printf("Error fetching block: %v", err)
		} else {
			fmt.Printf("Latest Block: %v\n", header.Number)
		}
		time.Sleep(15 * time.Second)
	}
}
