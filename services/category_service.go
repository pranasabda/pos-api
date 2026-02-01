package services

import (
	"KASIR-API/models"
	"KASIR-API/repositories"
)

type CategoryService interface {
	FindAll() ([]models.Category, error)
	Store(category models.Category) (models.Category, error)
	FindByID(id int) (models.Category, error)
	Update(id int, category models.Category) (models.Category, error) // WAJIB ADA DISINI
	Delete(id int) error                                              // WAJIB ADA DISINI
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *categoryService {
	return &categoryService{repo}
}

func (s *categoryService) FindAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) FindByID(id int) (models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) Store(category models.Category) (models.Category, error) {
	return s.repo.Create(category)
}

func (s *categoryService) Update(id int, category models.Category) (models.Category, error) {
	return s.repo.Update(id, category)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
