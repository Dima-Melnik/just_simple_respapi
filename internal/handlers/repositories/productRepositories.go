package repositories

import (
	"just_simple_api/internal/models"
	"log"

	"gorm.io/gorm"
)

type ProductRepositories interface {
	GetAllProducts(offSet int, limit int) ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	CreateProduct(data models.Product) error
	UpdateProduct(data models.Product, id int) error
	DeleteProduct(id int) error
}

type productRepositories struct {
	db *gorm.DB
}

func NewProductRepositories(db *gorm.DB) ProductRepositories {
	return productRepositories{db: db}
}

func (r productRepositories) GetAllProducts(offSet int, limit int) ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Offset(offSet).Limit(limit).Find(&products).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (r productRepositories) GetProductByID(id int) (models.Product, error) {
	var product models.Product

	if err := r.db.First(&product, id).Error; err != nil {
		log.Println(err)
		return product, err
	}

	return product, nil
}

func (r productRepositories) CreateProduct(data models.Product) error {
	if err := r.db.Create(&data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r productRepositories) UpdateProduct(data models.Product, id int) error {
	if err := r.db.Model(&models.Product{}).Where("id = ?", id).Updates(data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r productRepositories) DeleteProduct(id int) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
