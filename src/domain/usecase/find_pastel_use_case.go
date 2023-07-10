package usecase

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

type IFindPastelUseCase interface {
	Execute(pastel *model.Pastel)
}

type findPastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewFindPastelUseCase(repository dataprovider.IPastelRepository) IFindPastelUseCase {
	return &findPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *findPastelUseCase) Execute(pastel *model.Pastel) {
	fmt.Println("Find usecase")
}
