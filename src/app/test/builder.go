package test

import (
	"github.com/google/uuid"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

func BuildPastelWithIgredients(pastelName string, ingredientNames []string) model.Pastel {
	return model.Pastel{
		ID:          uuid.NewString(),
		Name:        pastelName,
		Price:       10.0,
		Ingredients: BuildIngredients(ingredientNames),
	}
}

func BuildIngredients(names []string) []model.Ingredient {
	ingredients := make([]model.Ingredient, 0)
	for _, name := range names {
		ingredients = append(ingredients, model.Ingredient{
			ID:   uuid.NewString(),
			Name: name,
			Desc: "300 gramas",
		})
	}
	return ingredients
}
