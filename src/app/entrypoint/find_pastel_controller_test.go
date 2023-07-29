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

func TestFindById_Success(t *testing.T) {
	db := gormdataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)
	pastelToSave := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repo.Create(pastelToSave)
	assert.Nil(t, err)

	req, err := http.NewRequest("GET", utils.BaseUri+"/"+pastelToSave.ID, nil)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repo).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusOK, resp.Code)

	var response model.FindByIdResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.True(t, response.Success)
	assert.EqualValues(t, pastelToSave.ID, response.Data.ID)
	assert.EqualValues(t, pastelToSave.Name, response.Data.Name)
	assert.EqualValues(t, pastelToSave.Price, response.Data.Price)
	assert.Len(t, response.Data.Ingredients, 1)
	assert.EqualValues(t, pastelToSave.Ingredients[0].ID, response.Data.Ingredients[0].ID)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Name, response.Data.Ingredients[0].Name)
	assert.EqualValues(t, pastelToSave.Ingredients[0].Desc, response.Data.Ingredients[0].Desc)

	assert.EqualValues(t, response.Links[0].Rel, "ingredients")
	assert.EqualValues(t, response.Links[0].Href, "localhost:3000/pasteis/"+response.Data.ID)
	assert.EqualValues(t, response.Links[0].Type, "DELETE")
}

func TestFindById_Notfound(t *testing.T) {
	db := gormdataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)

	req, err := http.NewRequest("GET", utils.BaseUri+"/1234", nil)
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
