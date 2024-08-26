package internal

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QuoteServiceSuite struct {
	suite.Suite
	service *QuoteService
	db      *sql.DB
}

func (suite *QuoteServiceSuite) SetupSuite() {
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
	repository := NewSQLiteQuoteRepository(db)
	request := NewAwesomeApi()
	ctx := context.Background()
	suite.service = NewQuoteService(ctx, repository, request)
}

func (suite *QuoteServiceSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *QuoteServiceSuite) TestGetQuote() {

	codeIn := "USD-BRL"

	_, err := suite.service.GetQuote(codeIn)
	suite.NoError(err)

}

func TestQuoteServiceSuite(t *testing.T) {
	suite.Run(t, new(QuoteServiceSuite))
}
