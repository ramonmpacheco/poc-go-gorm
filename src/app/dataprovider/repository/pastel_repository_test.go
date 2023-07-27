package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	dataerrors "github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/data_errors"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	domainerrors "github.com/ramonmpacheco/poc-go-gorm/domain/domain_errors"
	"github.com/ramonmpacheco/poc-go-gorm/domain/model"
	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreate_success(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)
	pastelToSave := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
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

	pastelToSave := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	assert.Nil(t, repository.Create(pastelToSave))

	var savedPastel entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel, "id = ?", pastelToSave.ID)

	pastelToSave2 := test.BuildPastelDomainWithIgredients("Pantaneiro 2", []string{"Carne"})
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
	pastelToSave := test.BuildPastelDomainWithIgredients("Boi Ralado", []string{"Carne", "Queijo", "Azeitona"})
	repository.Create(pastelToSave)
	var savedPastel entity.Pastel
	db.DB.Model(&entity.Pastel{}).
		Preload("Ingredients").
		First(&savedPastel, "id = ?", pastelToSave.ID)
	assert.Len(t, savedPastel.Ingredients, 3)
}

func TestCreate_error_constraint(t *testing.T) {
	mockDB := new(test.MockDB)
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{
		Error: gorm.ErrDuplicatedKey,
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
	assert.ErrorIs(t, err, dataerrors.ErrDuplicatedKey)
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

func TestFindById_Success(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)
	pastelToSave := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repository.Create(pastelToSave)
	assert.Nil(t, err)

	pastel, err := repository.FindById(pastelToSave.ID)

	assert.Nil(t, err)
	assert.EqualValues(t, pastelToSave.ID, pastel.ID)
	assert.EqualValues(t, pastelToSave.Name, pastel.Name)
	assert.EqualValues(t, pastelToSave.Price, pastel.Price)
	assert.Len(t, pastel.Ingredients, 1)
	assert.EqualValues(t, pastelToSave.Ingredients[0].ID, pastel.Ingredients[0].ID)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Name, pastel.Ingredients[0].Name)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Desc, pastel.Ingredients[0].Desc)
}

func TestFindById_Notfound(t *testing.T) {
	db := dataprovider.NewSqlite()
	repository := NewPastelRepository(db)

	pastel, err := repository.FindById("1234")

	assert.Nil(t, pastel)
	assert.NotNil(t, err)
	assert.ErrorIs(t, dataerrors.ErrNotFound, err)
}

func TestUpdate_Success(t *testing.T) {
	db := dataprovider.NewSqlite()
	repo := NewPastelRepository(db)
	pastel := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repo.Create(pastel)
	assert.Nil(t, err)

	pastel.Name = "Pantaneiro Update"
	pastel.Price = float32(12.50)
	pastel.Ingredients[0].Name = "Carne Update"
	pastel.Ingredients[0].Desc = "Update"

	err = repo.Update(pastel)
	assert.Nil(t, err)

	saved, _ := repo.FindById(pastel.ID)

	assert.EqualValues(t, pastel.Name, saved.Name)
	assert.EqualValues(t, pastel.Price, saved.Price)
	assert.EqualValues(t, pastel.Ingredients[0].Name, saved.Ingredients[0].Name)
	assert.EqualValues(t, pastel.Ingredients[0].Desc, saved.Ingredients[0].Desc)
	assert.EqualValues(t, pastel.Ingredients[0].Desc, saved.Ingredients[0].Desc)
	assert.NotNil(t, saved.Ingredients[0].CreatedAt)
	assert.NotNil(t, saved.Ingredients[0].UpdatedAt)
}

func TestUpdate_Error(t *testing.T) {
	db := dataprovider.NewSqlite()
	repo := NewPastelRepository(db)
	pastel := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})

	pastel.Name = "Pantaneiro Update"
	pastel.Price = float32(12.50)
	pastel.Ingredients[0].Name = "Carne Update"
	pastel.Ingredients[0].Desc = "Update"

	err := repo.Update(pastel)
	assert.NotNil(t, err)

	assert.EqualValues(t, "registro não encontrado", err.Error())
}
