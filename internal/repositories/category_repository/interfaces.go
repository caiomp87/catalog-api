package category_repository

import "github.com/caiomp87/catalog-api/internal/models"

type ICategory interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	CreateCategory(category *models.Category) error
}
