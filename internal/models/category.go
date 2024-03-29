package models

import (
	"time"
)

type Category struct {
	ID        int32     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}
