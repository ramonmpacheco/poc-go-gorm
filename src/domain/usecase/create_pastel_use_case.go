package usecase

import (
	"strings"
	"time"

	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/rs/xid"
)

type ICreatePastelUseCase interface {
	Create(pastel *model.Pastel) (string, error)
}

type createPastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewCreatePastelUseCase(repository dataprovider.IPastelRepository) ICreatePastelUseCase {
	return &createPastelUseCase{
		Repository: repository,
	}
}

func (cpuc *createPastelUseCase) Create(pastel *model.Pastel) (string, error) {
	pastel.ID = xid.New().String()
	pastel.Name = strings.ToUpper(pastel.Name)
	pastel.CreatedAt = time.Now()

	for i, v := range pastel.Ingredients {
		pastel.Ingredients[i].ID = xid.New().String()
		pastel.Ingredients[i].Name = strings.ToUpper(v.Name)
		pastel.Ingredients[i].CreatedAt = time.Now()
	}

	if err := cpuc.Repository.Create(*pastel); err != nil {
		return "", err
	}
	return pastel.ID, nil
}
