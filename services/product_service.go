package services

import (
	"kasir-api-go/models"
	"kasir-api-go/repositories"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(id uint, product *models.Product) error
	Delete(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) Update(id uint, product *models.Product) error {
	existingProduct, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Stock = product.Stock
	existingProduct.CategoryID = product.CategoryID

	return s.repo.Update(existingProduct)
}

func (s *productService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
