package usecase_test

import (
	"errors"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/ramonmpacheco/poc-go-gorm/domain/test"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePastel_success(t *testing.T) {
	repositoryMock := new(test.CreatePastelRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)

	usecase := usecase.NewCreatePastelUseCase(repositoryMock)
	id, err := usecase.Create(&model.Pastel{
		Name:  "Pantaneiro",
		Price: 10.0,
		Ingredients: []model.Ingredient{
			{
				Name: "Carne",
				Desc: "300 gramas",
			},
		},
	})

	assert.Nil(t, err)
	assert.Len(t, id, 36)
}

func TestCreatePastel_error(t *testing.T) {
	repositoryMock := new(test.CreatePastelRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("test error"))

	usecase := usecase.NewCreatePastelUseCase(repositoryMock)
	id, err := usecase.Create(&model.Pastel{})

	assert.EqualValues(t, "", id)
	assert.EqualValues(t, err.Error(), "test error")
}
