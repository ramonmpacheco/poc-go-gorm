package dataerrors

import (
	"errors"
	"testing"

	domainerrors "github.com/ramonmpacheco/poc-go-gorm/domain/domain_errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetProperError(t *testing.T) {
	assert.ErrorIs(t, GetProperError(errors.New("Internal")), domainerrors.ErrInternal)
	assert.ErrorIs(t, GetProperError(gorm.ErrDuplicatedKey), ErrDuplicatedKey)
}
