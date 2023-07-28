package usecase

import (
	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
)

type IDeletePastelUseCase interface {
	DeleteById(string) error
}

type deletePastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewDeletePastelUseCase(repository dataprovider.IPastelRepository) IFindPastelUseCase {
	return &findPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *findPastelUseCase) DeleteById(id string) error {
	return cpuc.Repository.DeleteById(id)
}
