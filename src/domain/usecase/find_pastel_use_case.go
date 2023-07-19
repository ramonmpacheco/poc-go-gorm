package usecase

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
)

type IFindPastelUseCase interface {
	FindById(string)
}

type findPastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewFindPastelUseCase(repository dataprovider.IPastelRepository) IFindPastelUseCase {
	return &findPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *findPastelUseCase) FindById(id string) {
	fmt.Println("Find usecase")
}
