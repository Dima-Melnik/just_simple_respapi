package db

import (
	"just_simple_api/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	connStr := ConnectString()

	DB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.ShopUser{}); err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}
}
