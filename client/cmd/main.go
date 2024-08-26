package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Quote struct {
	Bid float64 `json:"bid"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var quote Quote

	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		log.Fatal(err)
	}

	filename := "cotacao.txt"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("erro ao abrir ou criar o arquivo: %v", err)
	}
	defer file.Close()

	text := "Dólar: " + strconv.FormatFloat(quote.Bid, 'f', 2, 64) + "\n"

	if _, err := file.WriteString(text); err != nil {
		log.Fatalf("Erro ao escrever no arquivo: %v", err)
	}

}
