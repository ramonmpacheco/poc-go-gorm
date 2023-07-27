package usecase

import (
	"errors"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate_Success(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	pastelMock := test.BuildPastelDomainWithIgredients("Carne", []string{"Carne"})

	repositoryMock.On("Update", mock.Anything).Return(nil)
	repositoryMock.On("FindById", mock.Anything).Return(&pastelMock, nil)

	uc := NewUpdatePastelUseCase(repositoryMock)

	err := uc.Update("abc", pastelMock)
	assert.Nil(t, err)
}

func TestUpdate_Error(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	pastelMock := test.BuildPastelDomainWithIgredients("Carne", []string{"Carne"})

	repositoryMock.On("Update", mock.Anything).Return(errors.New("error test"))
	repositoryMock.On("FindById", mock.Anything).Return(&pastelMock, nil)

	uc := NewUpdatePastelUseCase(repositoryMock)

	err := uc.Update("abc", pastelMock)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error test", err.Error())
}
