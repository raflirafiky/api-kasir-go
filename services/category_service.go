package services

import (
	"kasir-api-go/models"
	"kasir-api-go/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Create(category *models.Category) error
	Update(id uint, category *models.Category) error
	Delete(id uint) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetByID(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) Update(id uint, category *models.Category) error {
	existingCategory, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	existingCategory.Name = category.Name
	existingCategory.Description = category.Description

	return s.repo.Update(existingCategory)
}

func (s *categoryService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
