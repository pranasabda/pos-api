package repositories

import (
	"KASIR-API/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error) // Baru
	Delete(id int) error                                           // Baru
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

// GetAll menggunakan Preload untuk Challenge Join (menampilkan data kategori)
func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	// Preload Category untuk fitur Join
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	// Proses insert yang memicu error foreign key jika CategoryID salah
	err := r.db.Create(&product).Error
	return product, err
}

// Fungsi Update Produk
func (r *productRepository) Update(id int, product models.Product) (models.Product, error) {
	var prod models.Product
	if err := r.db.First(&prod, id).Error; err != nil {
		return prod, err
	}
	// Menggunakan Updates untuk memperbarui data berdasarkan struct yang dikirim
	r.db.Model(&prod).Updates(product)
	return prod, nil
}

// Fungsi Delete Produk
func (r *productRepository) Delete(id int) error {
	return r.db.Delete(&models.Product{}, id).Error
}
