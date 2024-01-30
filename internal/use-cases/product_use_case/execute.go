package product_use_case

import (
	"github.com/caiomp87/catalog-api/internal/models"
	"github.com/caiomp87/catalog-api/internal/repositories/product_repository"
)

var ProductUseCase *productUseCase

type CreateProductUseCaseRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
	CategoryID  int32   `json:"categoryId"`
	Price       float64 `json:"price"`
}

type productUseCase struct {
	productRepository product_repository.IProduct
}

func NewProductUseCase(productRepository product_repository.IProduct) *productUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (c *productUseCase) FindAll() ([]*models.Product, error) {
	return c.productRepository.GetProducts()
}

func (c *productUseCase) FindByID(id string) (*models.Product, error) {
	return c.productRepository.GetProductByID(id)
}

func (c *productUseCase) FindByCategoryID(categoryID string) ([]*models.Product, error) {
	return c.productRepository.GetProductsByCategoryID(categoryID)
}

func (c *productUseCase) Create(request CreateProductUseCaseRequest) error {
	product := models.NewProduct(request.Name, request.Description, request.ImageURL, request.CategoryID, request.Price)
	return c.productRepository.CreateProduct(product)
}
