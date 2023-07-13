package converter

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	domain "github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

func ToPastelDomain(request model.CreatePastelRequest) *domain.Pastel {
	return &domain.Pastel{
		Name:        request.Name,
		Price:       request.Price,
		Ingredients: toIngredienteDomain(request.Ingredients),
	}
}

func toIngredienteDomain(ingredientes []model.CreateIngredientRequest) []domain.Ingredient {
	i := make([]domain.Ingredient, 0)
	for _, v := range ingredientes {
		i = append(i, domain.Ingredient{
			ID:   v.ID,
			Name: v.Name,
			Desc: v.Desc,
		})
	}
	return i
}
