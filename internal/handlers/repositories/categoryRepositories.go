package repositories

import (
	"just_simple_api/internal/models"
	"log"

	"gorm.io/gorm"
)

type CategoryRepositories interface {
	GetAllCategory() ([]models.Category, error)
	GetCategoryByID(id int) (models.Category, error)
	CreateCategory(data models.Category) error
	UpdateCategory(data models.Category, id int) error
	DeleteCategory(id int) error
}

type categoryRepositories struct {
	db *gorm.DB
}

func NewCategoryRepositories(db *gorm.DB) CategoryRepositories {
	return categoryRepositories{db: db}
}

func (r categoryRepositories) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return categories, nil
}

func (r categoryRepositories) GetCategoryByID(id int) (models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		log.Println(err)
		return category, err
	}

	return category, nil
}

func (r categoryRepositories) CreateCategory(data models.Category) error {
	if err := r.db.Create(&data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r categoryRepositories) UpdateCategory(data models.Category, id int) error {
	if err := r.db.Model(&models.Category{}).Where("id = ?", id).Updates(data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r categoryRepositories) DeleteCategory(id int) error {
	if err := r.db.Delete(&models.Category{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
