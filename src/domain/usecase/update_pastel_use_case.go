package usecase

import (
	"time"

	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

type IUpdatePastelUseCase interface {
	Update(string, model.Pastel) error
}

type updatePastelUseCase struct {
	Repository dataprovider.IPastelRepository
}

func NewUpdatePastelUseCase(repository dataprovider.IPastelRepository) IUpdatePastelUseCase {
	return &updatePastelUseCase{
		Repository: repository,
	}
}

func (cpuc *updatePastelUseCase) Update(id string, pastel model.Pastel) error {
	if _, err := cpuc.Repository.FindById(id); err != nil {
		return err
	}
	now := time.Now()
	pastel.ID = id
	pastel.CreatedAt = nil
	pastel.UpdatedAt = &now

	for i := range pastel.Ingredients {
		pastel.CreatedAt = nil
		pastel.Ingredients[i].UpdatedAt = &now
	}
	return cpuc.Repository.Update(pastel)
}
