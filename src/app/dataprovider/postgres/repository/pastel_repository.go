package repository

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/postgres/converter"
	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	domainerrors "github.com/ramonmpacheco/poc-go-gorm/domain/domain_errors"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"gorm.io/gorm"
)

type pastelRepository struct {
	Db *gorm.DB
}

func NewPastelRepository(db *gorm.DB) dataprovider.IPastelRepository {
	return &pastelRepository{
		Db: db,
	}
}

func (pr *pastelRepository) Create(pastel model.Pastel) error {
	// Need pass pointer in order to avoid error:
	// reflect.Value.Set using unaddressable value
	result := pr.Db.Create(converter.ToPastelEntity(pastel))
	if result.Error != nil {
		return domainerrors.ErrDatabaseSave
	}
	return nil
}
