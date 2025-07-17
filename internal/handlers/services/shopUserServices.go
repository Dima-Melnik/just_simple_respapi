package services

import (
	"just_simple_api/internal/handlers/repositories"
	"just_simple_api/internal/models"
)

type UserShopServices interface {
	GetAllUsers(offSet int, limit int) ([]models.ShopUser, error)
	GetUserByID(id int) (models.ShopUser, error)
	AddUser(data models.ShopUser) error
	UpdateUser(data models.ShopUser, id int) error
	DeleteUser(id int) error
}

type userShopServices struct {
	repo repositories.UserShopRepositories
}

func NewUserShopServices(r repositories.UserShopRepositories) UserShopServices {
	return userShopServices{repo: r}
}

func (s userShopServices) GetAllUsers(offSet int, limit int) ([]models.ShopUser, error) {
	return s.repo.GetAllUsers(offSet, limit)
}

func (s userShopServices) GetUserByID(id int) (models.ShopUser, error) {
	return s.repo.GetUserByID(id)
}

func (s userShopServices) AddUser(data models.ShopUser) error {
	return s.repo.AddUser(data)
}

func (s userShopServices) UpdateUser(data models.ShopUser, id int) error {
	return s.repo.UpdateUser(data, id)
}

func (s userShopServices) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
