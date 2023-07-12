package validator_test

import (
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/validator"

	"github.com/stretchr/testify/assert"
)

func TestValidateStruct_return_list_of_errors_when_invalid(t *testing.T) {
	carne := model.CreateIngredientRequest{}
	ingredientes := []model.CreateIngredientRequest{
		carne,
	}
	pastelCarne := model.CreatePastelRequest{
		Ingredients: ingredientes,
	}

	r, err := validator.ValidateStruct(pastelCarne)
	assert.EqualValues(t, "invalid data", err.Error())
	assert.EqualValues(t, "Name is required", r[0])
	assert.EqualValues(t, "Price is required", r[1])
	assert.EqualValues(t, "Name is required", r[2])
	assert.EqualValues(t, "Desc is required", r[3])
}

func TestValidateStruct_return_empty_errors_list_when_all_data_valid(t *testing.T) {
	carne := model.CreateIngredientRequest{
		Name: "Carne",
		Desc: "250 gramas",
	}
	ingredientes := []model.CreateIngredientRequest{
		carne,
	}
	pastelCarne := model.CreatePastelRequest{
		Name:        "Boi ralado",
		Price:       9.50,
		Ingredients: ingredientes,
	}

	r, err := validator.ValidateStruct(pastelCarne)
	assert.Nil(t, err)
	assert.Len(t, r, 0)
}
