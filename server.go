package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nero2009/grpcTrader/trader"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server is running on port :9092")

	s := trader.Server{}

	grpcServer := grpc.NewServer()

	trader.RegisterTraderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
