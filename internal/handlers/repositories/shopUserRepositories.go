package repositories

import (
	"just_simple_api/internal/models"
	"log"

	"gorm.io/gorm"
)

type UserShopRepositories interface {
	GetAllUsers(offSet int, limit int) ([]models.ShopUser, error)
	GetUserByID(id int) (models.ShopUser, error)
	AddUser(data models.ShopUser) error
	UpdateUser(data models.ShopUser, id int) error
	DeleteUser(id int) error
}

type userShopRepositories struct {
	db *gorm.DB
}

func NewUserShopRepositories(db *gorm.DB) UserShopRepositories {
	return userShopRepositories{db: db}
}

func (r userShopRepositories) GetAllUsers(offSet int, limit int) ([]models.ShopUser, error) {
	var users []models.ShopUser

	if err := r.db.Offset(offSet).Limit(limit).Find(&users).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (r userShopRepositories) GetUserByID(id int) (models.ShopUser, error) {
	var user models.ShopUser

	if err := r.db.First(&user, id).Error; err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r userShopRepositories) AddUser(id int) mode
