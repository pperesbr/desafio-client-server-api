package internal

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *QuoteService
}

func NewQuoteHandler(service *QuoteService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetQuote(w http.ResponseWriter, r *http.Request) {

	code := "USD-BRL"
	bid, err := h.service.GetQuote(code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bid)

}
