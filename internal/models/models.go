package models

import "time"

type Category struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShopUser struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
