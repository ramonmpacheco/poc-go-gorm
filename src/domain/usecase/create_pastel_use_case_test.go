package usecase

import (
	"errors"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePastel_success(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)

	uc := NewCreatePastelUseCase(repositoryMock)
	pastelMock := test.BuildPastelDomainWithIgredients("Carne", []string{"Carne"})
	id, err := uc.Create(&pastelMock)

	assert.Nil(t, err)
	assert.Len(t, id, 36)
}

func TestCreatePastel_error(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("test error"))

	uc := NewCreatePastelUseCase(repositoryMock)
	pastelMock := test.BuildEmptyPastelDomain()
	id, err := uc.Create(&pastelMock)

	assert.EqualValues(t, "", id)
	assert.EqualValues(t, err.Error(), "test error")
}
