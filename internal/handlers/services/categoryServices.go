package services

import (
	"just_simple_api/internal/handlers/repositories"
	"just_simple_api/internal/models"
)

type CategoryServices interface {
	GetAllCategory() ([]models.Category, error)
	GetCategoryByID(id int) (models.Category, error)
	CreateCategory(data models.Category) error
	UpdateCatagory(data models.Category, id int) error
	DeleteCategory(id int) error
}

type categoryServices struct {
	repo repositories.CategoryRepositories
}

func NewCategoryServices(r repositories.CategoryRepositories) CategoryServices {
	return &categoryServices{repo: r}
}

func (s categoryServices) GetAllCategory() ([]models.Category, error) {
	return s.repo.GetAllCategory()
}

func (s categoryServices) GetCategoryByID(id int) (models.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s categoryServices) CreateCategory(data models.Category) error {
	return s.repo.CreateCategory(data)
}

func (s categoryServices) UpdateCatagory(data models.Category, id int) error {
	return s.repo.UpdateCategory(data, id)
}

func (s categoryServices) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
