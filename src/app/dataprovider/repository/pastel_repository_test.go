package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	"github.com/ramonmpacheco/poc-go-gorm/app/test"
	domainerrors "github.com/ramonmpacheco/poc-go-gorm/domain/domain_errors"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreate_success(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)
	pastelToSave := test.BuildPastelWithIgredients("Pantaneiro", []string{"Carne"})
	err := repository.Create(pastelToSave)
	assert.Nil(t, err)

	var savedPastel entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel, "id = ?", pastelToSave.ID)

	assert.EqualValues(t, pastelToSave.ID, savedPastel.ID)
	assert.EqualValues(t, pastelToSave.Name, savedPastel.Name)
	assert.EqualValues(t, pastelToSave.Price, savedPastel.Price)
	assert.Len(t, savedPastel.Ingredients, 1)
	assert.EqualValues(t, pastelToSave.Ingredients[0].ID, savedPastel.Ingredients[0].ID)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Name, savedPastel.Ingredients[0].Name)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Desc, savedPastel.Ingredients[0].Desc)
}

func TestCreate_using_same_ingredient_twice(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)

	pastelToSave := test.BuildPastelWithIgredients("Pantaneiro", []string{"Carne"})
	assert.Nil(t, repository.Create(pastelToSave))

	var savedPastel entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel, "id = ?", pastelToSave.ID)

	pastelToSave2 := test.BuildPastelWithIgredients("Pantaneiro 2", []string{"Carne"})
	pastelToSave2.Ingredients[0].ID = pastelToSave.Ingredients[0].ID
	assert.Nil(t, repository.Create(pastelToSave2))

	var savedPastel2 entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel2, "id = ?", pastelToSave2.ID)

	var ingredients []entity.Ingredient
	db.DB.Find(&ingredients)

	assert.EqualValues(t, pastelToSave.ID, savedPastel.ID)
	assert.EqualValues(t, pastelToSave.Ingredients[0].ID, savedPastel.Ingredients[0].ID)
	assert.Len(t, ingredients, 1)
	assert.EqualValues(t, savedPastel2.ID, pastelToSave2.ID)
	assert.EqualValues(t, ingredients[0].ID, savedPastel2.Ingredients[0].ID)
}

func TestCreate_more_than_one_ingridient(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)
	pastelToSave := test.BuildPastelWithIgredients("Boi Ralado", []string{"Carne", "Queijo", "Azeitona"})
	repository.Create(pastelToSave)
	var savedPastel entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel, "id = ?", pastelToSave.ID)
	assert.Len(t, savedPastel.Ingredients, 3)
}

func TestCreate_internal_error(t *testing.T) {
	mockDB := new(test.MockDB)
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{
		Error: gorm.ErrInvalidDB,
	})

	repo := NewPastelRepository(mockDB)
	err := repo.Create(model.Pastel{
		ID:    uuid.NewString(),
		Name:  "Pantaneiro",
		Price: 10.0,
		Ingredients: []model.Ingredient{
			{
				ID:   uuid.NewString(),
				Name: "Carne",
				Desc: "300 gramas",
			},
		},
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domainerrors.ErrInternal)
}
