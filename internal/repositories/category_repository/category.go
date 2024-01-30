package category_repository

import (
	"database/sql"

	"github.com/caiomp87/catalog-api/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) ICategory {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) GetCategories() ([]*models.Category, error) {
	rows, err := cr.db.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*models.Category, 0)

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (cr *CategoryRepository) GetCategoryByID(id string) (*models.Category, error) {
	row := cr.db.QueryRow("SELECT id, name, created_at, updated_at FROM categories WHERE id = $1", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var category models.Category

	if err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *CategoryRepository) CreateCategory(category *models.Category) error {
	_, err := cr.db.Exec("INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}
