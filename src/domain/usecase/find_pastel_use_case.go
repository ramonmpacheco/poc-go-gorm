package usecase

import (
	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

type IFindPastelUseCase interface {
	FindById(string) (*model.Pastel, error)
}

type findPastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewFindPastelUseCase(repository dataprovider.IPastelRepository) IFindPastelUseCase {
	return &findPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *findPastelUseCase) FindById(id string) (*model.Pastel, error) {
	return cpuc.Repository.FindById(id)
}
