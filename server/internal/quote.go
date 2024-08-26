package internal

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	CodeIn    string    `json:"code_in"`
	Name      string    `json:"name"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	VarBid    float64   `json:"var_bid"`
	PctChange float64   `json:"pct_change"`
	Bid       float64   `json:"bid"`
	CreatedAt time.Time `json:"created_at"`
}

type QuoteRepository interface {
	Create(ctx context.Context, quote *Quote) error
}

type SQLiteQuoteRepository struct {
	db *sql.DB
}

func NewSQLiteQuoteRepository(db *sql.DB) *SQLiteQuoteRepository {
	return &SQLiteQuoteRepository{
		db: db,
	}
}

func (r *SQLiteQuoteRepository) Create(ctx context.Context, quote *Quote) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO quotes (id, code, code_in, name, high, low, var_bid, pct_change, bid, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		quote.ID,
		quote.Code,
		quote.CodeIn,
		quote.Name,
		quote.High,
		quote.Low,
		quote.VarBid,
		quote.PctChange,
		quote.Bid,
		quote.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

var _ QuoteRepository = &SQLiteQuoteRepository{}
