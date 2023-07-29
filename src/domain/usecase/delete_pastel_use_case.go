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

func NewDeletePastelUseCase(repository dataprovider.IPastelRepository) IDeletePastelUseCase {
	return &deletePastelUseCase{
		Repository: repository,
	}
}

func (cpuc *deletePastelUseCase) DeleteById(id string) error {
	return cpuc.Repository.DeleteById(id)
}
