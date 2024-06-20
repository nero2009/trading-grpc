package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/nero2009/grpcTrader/trader"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	http.HandleFunc("/toptraders", top10TradersHandler)
	http.ListenAndServe(":8083", nil)

	fmt.Print("Running leaderboard on port 9091")

}

func top10TradersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := getTopTradersFromGRPC()
	if err != nil {
		http.Error(w, "Failed to get trading info", http.StatusInternalServerError)
		return
	}

	sort.Slice(resp.Trades, func(i, j int) bool {
		return resp.Trades[i].Price > resp.Trades[j].Price
	})

	top10Trades := resp.Trades[:10]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(top10Trades)

}

func getTopTradersFromGRPC() (*trader.TradeResponse, error) {
	conn, err := grpc.NewClient("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := trader.NewTraderServiceClient(conn)

	resp, err := client.GetTradingInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to get trading info: %v", err)
	}

	return resp, err
}
