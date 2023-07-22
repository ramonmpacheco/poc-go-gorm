package test

import (
	"time"

	"github.com/google/uuid"
	entitymodel "github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	appmodel "github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

func BuildPastelDomainWithIgredients(pastelName string, ingredientNames []string) model.Pastel {
	return model.Pastel{
		ID:          uuid.NewString(),
		Name:        pastelName,
		Price:       10.0,
		Ingredients: BuildIngredientsDomain(ingredientNames),
	}
}

func BuildEmptyPastelDomain() model.Pastel {
	return model.Pastel{}
}

func BuildIngredientsDomain(names []string) []model.Ingredient {
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

func BuildCreatePastelRequest(name string, ingredientNames []string) appmodel.CreatePastelRequest {
	return appmodel.CreatePastelRequest{
		Name:        name,
		Price:       10.0,
		Ingredients: BuildCreateIngredientRequest(ingredientNames),
	}
}

func BuildCreateIngredientRequest(names []string) []appmodel.CreateIngredientRequest {
	ingredients := make([]appmodel.CreateIngredientRequest, 0)
	for _, name := range names {
		ingredients = append(ingredients, appmodel.CreateIngredientRequest{
			Name: name,
			Desc: "test",
		})
	}
	return ingredients
}

func BuildPastelEntityWithIgredients(pastelName string, ingredientNames []string) entitymodel.Pastel {
	now := time.Now()
	return entitymodel.Pastel{
		ID:          uuid.NewString(),
		Name:        pastelName,
		Price:       10.0,
		Ingredients: BuildIngredientsEntity(ingredientNames),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
}

func BuildIngredientsEntity(names []string) []entitymodel.Ingredient {
	now := time.Now()
	ingredients := make([]entitymodel.Ingredient, 0)
	for _, name := range names {
		ingredients = append(ingredients, entitymodel.Ingredient{
			ID:        uuid.NewString(),
			Name:      name,
			Desc:      "300 gramas",
			CreatedAt: &now,
			UpdatedAt: &now,
		})
	}
	return ingredients
}
