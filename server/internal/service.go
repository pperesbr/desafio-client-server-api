package internal

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type QuoteService struct {
	repo    QuoteRepository
	request Request
	ctx     context.Context
}

func NewQuoteService(ctx context.Context, repo QuoteRepository, request Request) *QuoteService {
	return &QuoteService{
		ctx:     ctx,
		repo:    repo,
		request: request,
	}
}

func (s *QuoteService) GetQuote(code string) (BidResponse, error) {

	ctxApi, cancelApi := context.WithTimeout(s.ctx, 200*time.Millisecond)
	defer cancelApi()

	quote, err := s.request.Do(ctxApi, code)

	if err != nil {
		return BidResponse{}, err
	}

	ctxRepository, cancelRepository := context.WithTimeout(s.ctx, 10*time.Millisecond)
	defer cancelRepository()

	quote.ID = uuid.New().String()
	err = s.repo.Create(ctxRepository, quote)

	if err != nil {
		return BidResponse{}, err
	}

	return BidResponse{Bid: quote.Bid}, nil

}
