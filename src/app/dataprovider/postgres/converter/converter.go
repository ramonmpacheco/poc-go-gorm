package converter

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/postgres/entity"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

func ToPastelEntity(pastel model.Pastel) *entity.Pastel {
	return &entity.Pastel{
		ID:          pastel.ID,
		Name:        pastel.Name,
		Price:       pastel.Price,
		Ingredients: toIngredienteEntity(pastel.Ingredients),
		CreatedAt:   &pastel.CreatedAt,
		UpdatedAt:   &pastel.UpdatedAt,
	}
}

func toIngredienteEntity(ingredientes []model.Ingredient) []entity.Ingredient {
	i := make([]entity.Ingredient, 0)
	for _, v := range ingredientes {
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
