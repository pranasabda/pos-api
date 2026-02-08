package services

import (
	"KASIR-API/models"
	"KASIR-API/repositories"
	"time"
)

type TransactionService interface {
	Checkout(req models.CheckoutRequest) (models.Transaction, error)
}

type transactionService struct {
	repo     repositories.TransactionRepository
	prodRepo repositories.ProductRepository
}

func NewTransactionService(r repositories.TransactionRepository, pr repositories.ProductRepository) *transactionService {
	return &transactionService{r, pr}
}

func (s *transactionService) Checkout(req models.CheckoutRequest) (models.Transaction, error) {
	var transaction models.Transaction
	var totalAmount int

	for _, item := range req.Items {
		// Validasi produk dan ambil harga asli dari database
		product, err := s.prodRepo.GetByID(item.ProductID)
		if err != nil {
			return transaction, err
		}

		subtotal := product.Price * item.Quantity
		totalAmount += subtotal

		detail := models.TransactionDetail{
			ProductID:   int(product.ID),
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		}
		transaction.Details = append(transaction.Details, detail)
	}

	transaction.TotalAmount = totalAmount
	transaction.CreatedAt = time.Now()

	return s.repo.Create(transaction)
}
