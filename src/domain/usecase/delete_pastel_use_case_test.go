package usecase

import (
	"errors"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteById_Success(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	repositoryMock.On("DeleteById", mock.Anything).Return(nil)

	uc := NewDeletePastelUseCase(repositoryMock)
	err := uc.DeleteById("abc")

	assert.Nil(t, err)
}

func TestDeleteById_Not_Found(t *testing.T) {
	repositoryMock := new(test.PastelRepositoryMock)
	repositoryMock.On("DeleteById", mock.Anything).Return(errors.New("not found"))

	uc := NewDeletePastelUseCase(repositoryMock)
	err := uc.DeleteById("abc")

	assert.EqualValues(t, "not found", err.Error())
}
