package dataerrors

import (
	"errors"

	domainerrors "github.com/ramonmpacheco/poc-go-gorm/domain/domain_errors"
	"gorm.io/gorm"
)

var ErrNotFound error = errors.New("registro não encontrado")
var ErrDuplicatedKey error = errors.New("registro já existente, verifique os dados enviados")

func GetProperError(err error) error {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrDuplicatedKey
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return domainerrors.ErrInternal
}
