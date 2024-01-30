package product_repository

import (
	"database/sql"

	"github.com/caiomp87/catalog-api/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProduct {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]*models.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, description, price, category_id, image_url, created_at, updated_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*models.Product, 0)

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (pr *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	row := pr.db.QueryRow("SELECT id, name, description, price, category_id, image_url, created_at, updated_at FROM products WHERE id = $1", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var product models.Product

	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) GetProductsByCategoryID(categoryID string) ([]*models.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, description, price, category_id, image_url,created_at, updated_at FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*models.Product, 0)

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	_, err := pr.db.Exec("INSERT INTO products (name, description, price, category_id, image_url) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return err
	}

	return nil
}
