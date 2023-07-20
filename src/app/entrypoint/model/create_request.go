package model

type CreatePastelRequest struct {
	Name        string                    `json:"name" validate:"required"`
	Price       float32                   `json:"price" validate:"required"`
	Ingredients []CreateIngredientRequest `json:"ingredients" validate:"min=1,dive"`
}

type CreateIngredientRequest struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc" validate:"required"`
}
