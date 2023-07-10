package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/rs/xid"
)

type ICreatePastelUseCase interface {
	Execute(pastel *model.Pastel)
}

type createPastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewCreatePastelUseCase(repository dataprovider.IPastelRepository) ICreatePastelUseCase {
	return &createPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *createPastelUseCase) Execute(pastel *model.Pastel) {
	pastel.ID = xid.New().String()
	pastel.Name = strings.ToUpper(pastel.Name)
	pastel.CreatedAt = time.Now()

	for i, v := range pastel.Ingredients {
		pastel.Ingredients[i].ID = xid.New().String()
		pastel.Ingredients[i].Name = strings.ToUpper(v.Name)
		pastel.Ingredients[i].CreatedAt = time.Now()
	}

	if err := cpuc.Repository.Create(*pastel); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Pastel created successfully 😋")
}
