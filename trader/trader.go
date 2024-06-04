package trader

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
}

func (s *Server) GetTradingInfo(ctx context.Context, in *emptypb.Empty) (*TradeResponse, error) {
	var tradingInfo = make([]*Trade, 0)
	var cryptoSymbols = []string{"BTC", "ETH", "LTC", "XRP", "BCH", "EOS", "XLM", "ADA", "TRX", "XMR"}

	for i := 0; i < 1000; i++ {
		tradingInfo = append(tradingInfo, &Trade{
			StockSymbol: cryptoSymbols[gofakeit.Number(0, len(cryptoSymbols)-1)],
			Price:       gofakeit.Price(1, 1000),
			LastUpdated: time.Now().Unix(),
			TraderName:  gofakeit.Name(),
		})
	}

	return &TradeResponse{
		Trades: tradingInfo,
	}, nil

}

func (s *Server) mustEmbedUnimplementedTraderServiceServer() {}
