package models

import (
	"time"
)

type Product struct {
	ID          int32     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	CategoryID  int32     `json:"categoryId" db:"category_id"`
	ImageURL    string    `json:"imageUrl" db:"image_url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

func NewProduct(name, description, imageURL string, categoryID int32, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
