package test

import (
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/stretchr/testify/mock"
)

type CreatePastelRepositoryMock struct {
	mock.Mock
}

func (c *CreatePastelRepositoryMock) Create(pastel model.Pastel) error {
	args := c.Called(pastel)
	return args.Error(0)
}
