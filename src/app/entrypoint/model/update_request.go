package model

type UpdatePastelRequest struct {
	Name        string                    `json:"name" validate:"required"`
	Price       float32                   `json:"price" validate:"required"`
	Ingredients []UpdateIngredientRequest `json:"ingredients" validate:"min=1,dive"`
}

type UpdateIngredientRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc" validate:"required"`
}
