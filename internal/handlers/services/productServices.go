package services

import (
	"just_simple_api/internal/handlers/repositories"
	"just_simple_api/internal/models"
)

type ProductServices interface {
	GetAllProducts(offSet int, limit int) ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	CreateProduct(data models.Product) error
	UpdateProduct(data models.Product, id int) error
	DeleteProduct(id int) error
}

type productServices struct {
	repo repositories.ProductRepositories
}

func NewProductRepositories(r repositories.ProductRepositories) ProductServices {
	return productServices{repo: r}
}

func (s productServices) GetAllProducts(offSet int, limit int) ([]models.Product, error) {
	return s.repo.GetAllProducts(offSet, limit)
}

func (s productServices) GetProductByID(id int) (models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s productServices) CreateProduct(data models.Product) error {
	return s.repo.CreateProduct(data)
}

func (s productServices) UpdateProduct(data models.Product, id int) error {
	return s.repo.UpdateProduct(data, id)
}

func (s productServices) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
