package product_repository

import "github.com/caiomp87/catalog-api/internal/models"

type IProduct interface {
	GetProducts() ([]*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	GetProductsByCategoryID(categoryID string) ([]*models.Product, error)
	CreateProduct(product *models.Product) error
}
