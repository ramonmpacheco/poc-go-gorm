package converter

import (
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/app/test"
	"github.com/stretchr/testify/assert"
)

func TestConvertToPastelEntity(t *testing.T) {
	pastelDomain := test.BuildPastelWithIgredients("4 Queijos", []string{"Mussarela", "Catupiry", "Parmesão", "Provolone"})
	pastelEntity := ToPastelEntity(pastelDomain)

	assert.EqualValues(t, pastelDomain.ID, pastelEntity.ID)
	assert.EqualValues(t, pastelDomain.Name, pastelEntity.Name)
	assert.EqualValues(t, pastelDomain.Price, pastelEntity.Price)

	assert.EqualValues(t, pastelDomain.Ingredients[0].ID, pastelEntity.Ingredients[0].ID)
	assert.EqualValues(t, pastelDomain.Ingredients[0].Name, pastelEntity.Ingredients[0].Name)
	assert.EqualValues(t, pastelDomain.Ingredients[0].Desc, pastelEntity.Ingredients[0].Desc)

	assert.EqualValues(t, pastelDomain.Ingredients[1].ID, pastelEntity.Ingredients[1].ID)
	assert.EqualValues(t, pastelDomain.Ingredients[1].Name, pastelEntity.Ingredients[1].Name)
	assert.EqualValues(t, pastelDomain.Ingredients[1].Desc, pastelEntity.Ingredients[1].Desc)

	assert.EqualValues(t, pastelDomain.Ingredients[2].ID, pastelEntity.Ingredients[2].ID)
	assert.EqualValues(t, pastelDomain.Ingredients[2].Name, pastelEntity.Ingredients[2].Name)
	assert.EqualValues(t, pastelDomain.Ingredients[2].Desc, pastelEntity.Ingredients[2].Desc)
}
