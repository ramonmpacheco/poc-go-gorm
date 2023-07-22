package converter

import (
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
)

func TestToPastelDomain_pastel_and_ingredients_success(t *testing.T) {
	pastelRequest := test.BuildCreatePastelRequest("Carne", []string{"Carne", "Azeitona"})

	pastelDomain := ToPastelDomainFromCreate(pastelRequest)

	assert.EqualValues(t, pastelRequest.Name, pastelDomain.Name)
	assert.EqualValues(t, pastelRequest.Price, pastelDomain.Price)
	assert.Len(t, pastelDomain.Ingredients, 2)
	assert.EqualValues(t, pastelRequest.Ingredients[0].Name, pastelDomain.Ingredients[0].Name)
	assert.EqualValues(t, pastelRequest.Ingredients[0].Desc, pastelDomain.Ingredients[0].Desc)
	assert.EqualValues(t, pastelRequest.Ingredients[1].Name, pastelDomain.Ingredients[1].Name)
	assert.EqualValues(t, pastelRequest.Ingredients[1].Desc, pastelDomain.Ingredients[1].Desc)

}

func TestToPastelResponse(t *testing.T) {
	pastelDomain := test.BuildPastelDomainWithIgredients("Carne", []string{"Carne", "Azeitona"})

	response := ToPastelResponse(pastelDomain)

	assert.EqualValues(t, response.Name, pastelDomain.Name)
	assert.EqualValues(t, response.Price, pastelDomain.Price)
	assert.Len(t, pastelDomain.Ingredients, 2)
	assert.EqualValues(t, response.Ingredients[0].Name, pastelDomain.Ingredients[0].Name)
	assert.EqualValues(t, response.Ingredients[0].Desc, pastelDomain.Ingredients[0].Desc)
	assert.EqualValues(t, response.Ingredients[1].Name, pastelDomain.Ingredients[1].Name)
	assert.EqualValues(t, response.Ingredients[1].Desc, pastelDomain.Ingredients[1].Desc)
}
