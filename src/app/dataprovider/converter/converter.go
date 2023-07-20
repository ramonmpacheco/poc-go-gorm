package converter

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

func ToPastelEntity(pastel model.Pastel) *entity.Pastel {
	return &entity.Pastel{
		ID:          pastel.ID,
		Name:        pastel.Name,
		Price:       pastel.Price,
		Ingredients: toIngredientEntity(pastel.Ingredients),
		CreatedAt:   &pastel.CreatedAt,
		UpdatedAt:   &pastel.UpdatedAt,
	}
}

func toIngredientEntity(ingredients []model.Ingredient) []entity.Ingredient {
	i := make([]entity.Ingredient, 0)
	for _, v := range ingredients {
		i = append(i, entity.Ingredient{
			ID:        v.ID,
			Name:      v.Name,
			Desc:      v.Desc,
			CreatedAt: &v.CreatedAt,
			UpdatedAt: &v.UpdatedAt,
		})
	}
	return i
}

func ToPastelDomain(pastel entity.Pastel) *model.Pastel {
	return &model.Pastel{
		ID:          pastel.ID,
		Name:        pastel.Name,
		Price:       pastel.Price,
		Ingredients: toIngredientDomain(pastel.Ingredients),
		CreatedAt:   *pastel.CreatedAt,
		UpdatedAt:   *pastel.UpdatedAt,
	}
}

func toIngredientDomain(ingredients []entity.Ingredient) []model.Ingredient {
	i := make([]model.Ingredient, 0)
	for _, v := range ingredients {
		i = append(i, model.Ingredient{
			ID:        v.ID,
			Name:      v.Name,
			Desc:      v.Desc,
			CreatedAt: *v.CreatedAt,
			UpdatedAt: *v.UpdatedAt,
		})
	}
	return i
}
