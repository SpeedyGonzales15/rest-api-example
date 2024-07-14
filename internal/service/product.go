package service

import (
	"rest-api-example/internal/models"
	"rest-api-example/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductsById(id int) (models.Product, error) {
	return s.repo.GetProductsById(id)
}
