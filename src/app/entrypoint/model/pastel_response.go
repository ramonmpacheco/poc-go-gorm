package model

import "time"

type PastelResponse struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Price       float32              `json:"price"`
	Ingredients []IngredientResponse `json:"ingredients"`
	CreatedAt   *time.Time           `json:"created_at"`
	UpdatedAt   *time.Time           `json:"updated_at"`
}

type IngredientResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
