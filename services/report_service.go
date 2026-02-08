package services

import (
	"KASIR-API/repositories"
)

type ReportService interface {
	GetSummary(startDate, endDate string) (map[string]interface{}, error)
}

type reportService struct {
	repo repositories.TransactionRepository
}

func NewReportService(repo repositories.TransactionRepository) *reportService {
	return &reportService{repo}
}

func (s *reportService) GetSummary(startDate, endDate string) (map[string]interface{}, error) {
	transactions, err := s.repo.GetReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	var totalRevenue int
	productSales := make(map[string]int)

	for _, t := range transactions {
		totalRevenue += t.TotalAmount
		for _, d := range t.Details {
			productSales[d.ProductName] += d.Quantity
		}
	}

	bestSeller := "-"
	maxQty := 0
	for name, qty := range productSales {
		if qty > maxQty {
			maxQty = qty
			bestSeller = name
		}
	}

	return map[string]interface{}{
		"total_revenue":   totalRevenue,
		"total_transaksi": len(transactions),
		"produk_terlaris": map[string]interface{}{
			"nama":        bestSeller,
			"qty_terjual": maxQty,
		},
	}, nil
}
