package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pperesbr/desafio-client-server-api/server/internal"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")

	defer db.Close()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	service := internal.NewQuoteService(ctx, internal.NewSQLiteQuoteRepository(db), internal.NewAwesomeApi())
	handler := internal.NewQuoteHandler(service)

	http.HandleFunc("/", handler.GetQuote)
	http.ListenAndServe(":8080", nil)

}
