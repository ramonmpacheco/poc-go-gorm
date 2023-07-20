package dataprovider

import "github.com/ramonmpacheco/poc-go-gorm/domain/model"

type IPastelRepository interface {
	Create(pastel model.Pastel) error
	FindById(id string) (*model.Pastel, error)
}
