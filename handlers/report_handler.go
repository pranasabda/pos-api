package handlers

import (
	"KASIR-API/services"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ReportHandler struct {
	service services.ReportService
}

func NewReportHandler(s services.ReportService) *ReportHandler {
	return &ReportHandler{s}
}

func (h *ReportHandler) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	today := time.Now().Format("2006-01-02")
	var start, end string

	// Pengecekan apakah request datang dari endpoint hari-ini atau filter tanggal
	if strings.Contains(r.URL.Path, "hari-ini") {
		start, end = today, today
	} else {
		start = r.URL.Query().Get("start_date")
		end = r.URL.Query().Get("end_date")
		// Default ke hari ini jika parameter kosong
		if start == "" {
			start, end = today, today
		}
	}

	res, err := h.service.GetSummary(start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(res)
}
