package category_use_case

import (
	"github.com/caiomp87/catalog-api/internal/models"
	"github.com/caiomp87/catalog-api/internal/repositories/category_repository"
)

var CategoryUseCase *categoryUseCase

type CreateCategoryUseCaseRequest struct {
	Name string `json:"name"`
}

type categoryUseCase struct {
	categoryRepository category_repository.ICategory
}

func NewCategoryUseCase(categoryRepository category_repository.ICategory) *categoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (c *categoryUseCase) FindAll() ([]*models.Category, error) {
	return c.categoryRepository.GetCategories()
}

func (c *categoryUseCase) FindByID(id string) (*models.Category, error) {
	return c.categoryRepository.GetCategoryByID(id)
}

func (c *categoryUseCase) Create(request CreateCategoryUseCaseRequest) error {
	category := models.NewCategory(request.Name)
	return c.categoryRepository.CreateCategory(category)
}
