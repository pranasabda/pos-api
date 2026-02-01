package services

import (
	"KASIR-API/models"
	"KASIR-API/repositories"
)

type ProductService interface {
	FindAll() ([]models.Product, error)
	FindByID(id int) (models.Product, error)
	Store(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error) // WAJIB ADA DISINI
	Delete(id int) error                                           // WAJIB ADA DISINI
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *productService {
	return &productService{repo}
}

func (s *productService) FindAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) FindByID(id int) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) Store(product models.Product) (models.Product, error) {
	return s.repo.Create(product)
}

func (s *productService) Update(id int, product models.Product) (models.Product, error) {
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}
