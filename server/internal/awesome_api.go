package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ResponseAwesomeApi struct {
	Code      string `json:"code"`
	CodeIn    string `json:"codein"`
	Name      string `json:"name"`
	High      string `json:"high"`
	Low       string `json:"low"`
	VarBid    string `json:"varBid"`
	PctChange string `json:"pctChange"`
	Bid       string `json:"bid"`
	CreatedAt string `json:"create_date"`
}

type AwesomeApi struct {
}

// Do implements Request.
func (a *AwesomeApi) Do(ctx context.Context, code string) (*Quote, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s", code), nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response map[string]ResponseAwesomeApi

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	var quote Quote
	quote.Code = response[strings.ReplaceAll(code, "-", "")].Code
	quote.CodeIn = response[strings.ReplaceAll(code, "-", "")].CodeIn
	quote.Name = response[strings.ReplaceAll(code, "-", "")].Name

	quote.High, err = strconv.ParseFloat(response[strings.ReplaceAll(code, "-", "")].High, 64)

	if err != nil {
		return nil, err
	}

	quote.Low, err = strconv.ParseFloat(response[strings.ReplaceAll(code, "-", "")].Low, 64)

	if err != nil {
		return nil, err
	}

	quote.VarBid, err = strconv.ParseFloat(response[strings.ReplaceAll(code, "-", "")].VarBid, 64)

	if err != nil {
		fmt.Println("entrou aqui")
		return nil, err
	}

	quote.PctChange, err = strconv.ParseFloat(response[strings.ReplaceAll(code, "-", "")].PctChange, 64)

	if err != nil {
		return nil, err
	}

	quote.Bid, err = strconv.ParseFloat(response[strings.ReplaceAll(code, "-", "")].Bid, 64)

	if err != nil {
		return nil, err
	}

	layout := "2006-01-02 15:04:05"

	quote.CreatedAt, err = time.Parse(layout, response[strings.ReplaceAll(code, "-", "")].CreatedAt)
	if err != nil {
		return nil, err
	}

	return &quote, nil

}

func NewAwesomeApi() *AwesomeApi {
	return &AwesomeApi{}
}

var _ Request = &AwesomeApi{}
