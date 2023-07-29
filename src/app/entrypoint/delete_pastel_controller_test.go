package entrypoint_test

import (
	"encoding/json"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/gorm_internal/gorm_dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/app/rest"
	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/ramonmpacheco/poc-go-gorm/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteById_Success(t *testing.T) {
	db := gormdataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)
	pastel := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repo.Create(pastel)
	assert.Nil(t, err)

	req, err := http.NewRequest("DELETE", utils.BaseUri+"/"+pastel.ID, nil)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repo).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusNoContent, resp.Code)

	var response model.DeleteResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.EqualValues(t, "resource successfuly deleted", response.Message)
	assert.Len(t, response.Links, 0)

	_, err = repo.FindById(pastel.ID)
	assert.NotEqualValues(t, "not found", err.Error())
}

func TestDeleteById_When_Two_And_Delete_One(t *testing.T) {
	db := gormdataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)

	pastel1 := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repo.Create(pastel1)
	assert.Nil(t, err)

	pastel2 := test.BuildPastelDomainWithIgredients("Boi Ralado", []string{})
	pastel2.Ingredients = pastel1.Ingredients
	err = repo.Create(pastel2)
	assert.Nil(t, err)

	req, err := http.NewRequest("DELETE", utils.BaseUri+"/"+pastel1.ID, nil)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repo).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusNoContent, resp.Code)

	var response model.DeleteResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.EqualValues(t, "resource successfuly deleted", response.Message)
	assert.Len(t, response.Links, 0)

	saved, err := repo.FindById(pastel2.ID)
	assert.Nil(t, err)

	assert.EqualValues(t, pastel2.Name, saved.Name)
	assert.EqualValues(t, pastel2.Price, saved.Price)
	assert.EqualValues(t, pastel2.Ingredients[0].Name, saved.Ingredients[0].Name)
	assert.EqualValues(t, pastel2.Ingredients[0].Desc, saved.Ingredients[0].Desc)
	assert.EqualValues(t, pastel2.Ingredients[0].Desc, saved.Ingredients[0].Desc)
	assert.NotNil(t, saved.Ingredients[0].CreatedAt)
	assert.NotNil(t, saved.Ingredients[0].UpdatedAt)
}

func TestDeleteById_Not_Found(t *testing.T) {
	db := gormdataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)

	req, err := http.NewRequest("DELETE", utils.BaseUri+"/abc", nil)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repo).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusNotFound, resp.Code)

	var response model.ErrorResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.NotNil(t, response)
	assert.False(t, response.Success)
	assert.EqualValues(t, "registro não encontrado", response.Message)
	assert.Empty(t, response.Errors)
}
