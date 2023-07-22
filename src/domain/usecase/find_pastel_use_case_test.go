package usecase

import (
	"errors"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindById_Success(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	pastelMock := test.BuildPastelDomainWithIgredients("Carne", []string{"Carne"})
	repositoryMock.On("FindById", mock.Anything).Return(&pastelMock, nil)

	uc := NewFindPastelUseCase(repositoryMock)
	pastel, err := uc.FindById("abc")
	assert.Nil(t, err)

	assert.EqualValues(t, pastelMock.ID, pastel.ID)
	assert.EqualValues(t, pastelMock.Name, pastel.Name)
	assert.EqualValues(t, pastelMock.Price, pastel.Price)
	assert.Len(t, pastel.Ingredients, 1)
	assert.EqualValues(t, pastelMock.Ingredients[0].ID, pastel.Ingredients[0].ID)
	assert.EqualValues(t, pastelMock.Ingredients[0].Name, pastel.Ingredients[0].Name)
	assert.EqualValues(t, pastelMock.Ingredients[0].Desc, pastel.Ingredients[0].Desc)
}

func TestFindById_Error(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)

	repositoryMock.On("FindById", mock.Anything).Return(nil, errors.New("test error"))

	uc := NewFindPastelUseCase(repositoryMock)
	pastel, err := uc.FindById("abc")

	assert.Nil(t, pastel)
	assert.EqualValues(t, err.Error(), "test error")
}
