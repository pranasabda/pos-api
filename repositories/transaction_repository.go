package repositories

import (
	"KASIR-API/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction models.Transaction) (models.Transaction, error)
	GetReport(startDate, endDate string) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	// GORM akan otomatis memasukkan data ke tabel 'transactions' dan 'transaction_details'
	// karena sudah mendefinisikan relasi di struct. Ini menggantikan loop tx pada catatan session ke 3
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) GetReport(startDate, endDate string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	// Preload "Details" agar data item transaksi muncul di laporan
	err := r.db.Preload("Details").
		Where("created_at BETWEEN ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59").
		Find(&transactions).Error
	return transactions, err
}
