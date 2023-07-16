package repository

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/converter"
	dataerrors "github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/data_errors"
	domaindataprovider "github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
)

type pastelRepository struct {
	Db dataprovider.IDatabase
}

func NewPastelRepository(db dataprovider.IDatabase) domaindataprovider.IPastelRepository {
	return &pastelRepository{
		Db: db,
	}
}

func (pr *pastelRepository) Create(pastel model.Pastel) error {
	// Need pass pointer in order to avoid error:
	// reflect.Value.Set using unaddressable value
	result := pr.Db.Create(converter.ToPastelEntity(pastel))
	if result.Error != nil {
		return dataerrors.GetProperError(result.Error)
	}
	return nil
}
