package handlers

import (
	"KASIR-API/models"
	"KASIR-API/services"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(s services.TransactionService) *TransactionHandler {
	return &TransactionHandler{s}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest

	// Decode JSON Body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	// Proses Checkout
	res, err := h.service.Checkout(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
