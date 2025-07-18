package models

import "time"

type Category struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShopUser struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
