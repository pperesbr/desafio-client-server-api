package internal

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteQuoteRepositorySuite struct {
	suite.Suite
	repo *SQLiteQuoteRepository
	db   *sql.DB
}

func (suite *SQLiteQuoteRepositorySuite) SetupSuite() {

	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(suite.T(), err)

	_, err = db.Exec(`
		CREATE TABLE quotes (
			id TEXT PRIMARY KEY,
			code TEXT,
			code_in TEXT,
			name TEXT,
			high REAL,
			low REAL,
			var_bid REAL,
			pct_change REAL,
			bid REAL,
			created_at DATETIME
		)
	`)
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.repo = NewSQLiteQuoteRepository(db)
}

func (suite *SQLiteQuoteRepositorySuite) TearDownSuite() {
	// Close the database connection
	suite.db.Close()
}

func (suite *SQLiteQuoteRepositorySuite) TestCreate() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	quote := &Quote{
		ID:        uuid.NewString(),
		Code:      "USD",
		CodeIn:    "BRL",
		Name:      "Dólar Americano/Real Brasileiro",
		High:      5.6061,
		Low:       5.4738,
		VarBid:    0.0006,
		PctChange: 0.01,
		Bid:       5.4857,
		CreatedAt: time.Now(),
	}

	err := suite.repo.Create(ctx, quote)
	suite.NoError(err)

	row := suite.db.QueryRow("SELECT id, code, code_in, name, high, low, var_bid, pct_change, bid, created_at FROM quotes WHERE id = ?", quote.ID)
	var (
		id        string
		code      string
		codeIn    string
		name      string
		high      float64
		low       float64
		varBid    float64
		pctChange float64
		bid       float64
		createdAt time.Time
	)

	err = row.Scan(&id, &code, &codeIn, &name, &high, &low, &varBid, &pctChange, &bid, &createdAt)
	suite.NoError(err)

	suite.Equal(quote.ID, id)
	suite.Equal(quote.Code, code)
	suite.Equal(quote.CodeIn, codeIn)
	suite.Equal(quote.Name, name)
	suite.Equal(quote.High, high)
	suite.Equal(quote.Low, low)
	suite.Equal(quote.VarBid, varBid)
	suite.Equal(quote.PctChange, pctChange)
	suite.Equal(quote.Bid, bid)
	suite.Equal(quote.CreatedAt.Format(time.RFC3339), createdAt.Format(time.RFC3339))

	quote.ID = uuid.NewString()
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel2()
	err = suite.repo.Create(ctx2, quote)
	suite.ErrorIs(err, context.DeadlineExceeded)

}

func TestSQLiteQuoteRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteQuoteRepositorySuite))
}
