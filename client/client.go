package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nero2009/grpcTrader/trader"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := trader.NewTraderServiceClient(conn)

	resp, err := client.GetTradingInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to get trading info: %v", err)
	}

	fmt.Println("Trading Info:")
	for _, trade := range resp.Trades {
		fmt.Printf("Stock Symbol: %s, Price: %f, Last Updated: %d, Trader Name: %s\n", trade.StockSymbol, trade.Price, trade.LastUpdated, trade.TraderName)
	}
}
