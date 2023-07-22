package test

import (
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/stretchr/testify/mock"
)

type PastelRepositoryMock struct {
	mock.Mock
}

func (prm *PastelRepositoryMock) Create(pastel model.Pastel) error {
	args := prm.Called(pastel)
	return args.Error(0)
}

func (prm *PastelRepositoryMock) FindById(id string) (*model.Pastel, error) {
	args := prm.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Pastel), nil
}

func (prm *PastelRepositoryMock) Update(pastel model.Pastel) error {
	args := prm.Called(pastel)
	return args.Error(0)
}
