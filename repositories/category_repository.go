package repositories

import (
	"KASIR-API/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)                               // Get All Kategori
	GetByID(id int) (models.Category, error)                          // Get By ID Kategori
	Create(category models.Category) (models.Category, error)         // Create Kategori
	Update(id int, category models.Category) (models.Category, error) // Update Kategori
	Delete(id int) error                                              // Delete Kategori
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Create(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) GetByID(id int) (models.Category, error) {
	var category models.Category
	// err := r.db.Preload("Category").First(&category, id).Error
	err := r.db.First(&category, id).Error
	return category, err
}

// Fungsi Update
func (r *categoryRepository) Update(id int, category models.Category) (models.Category, error) {
	var cat models.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return cat, err
	}
	// Update kolom name dan description
	r.db.Model(&cat).Updates(category)
	return cat, nil
}

// Fungsi Delete
func (r *categoryRepository) Delete(id int) error {
	return r.db.Delete(&models.Category{}, id).Error
}
