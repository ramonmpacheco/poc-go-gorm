package model

import "time"

type Pastel struct {
	ID          string
	Name        string
	Price       float32
	Ingredients []Ingredient
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
