package entrypoint_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/app/rest"
	"github.com/ramonmpacheco/poc-go-gorm/test"
	"github.com/ramonmpacheco/poc-go-gorm/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_Success(t *testing.T) {
	db := dataprovider.NewSqlite()
	repo := repository.NewPastelRepository(db)
	pastel := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})
	err := repo.Create(pastel)
	assert.Nil(t, err)

	var buf bytes.Buffer
	pastel.Name = "Pantaneiro Update"
	pastel.Price = float32(12.50)
	pastel.Ingredients[0].Name = "Carne Update"
	pastel.Ingredients[0].Desc = "Update"
	json.NewEncoder(&buf).Encode(
		pastel,
	)

	req, err := http.NewRequest("PUT", utils.BaseUri+"/"+pastel.ID, &buf)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repo).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusOK, resp.Code)

	var response model.UpdateResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.True(t, response.Success)
	assert.EqualValues(t, "resource successfuly updated", response.Message)

	assert.EqualValues(t, response.Links[0].Rel, "ingredients")
	assert.EqualValues(t, response.Links[0].Href, "localhost:3000/pasteis/"+pastel.ID)
	assert.EqualValues(t, response.Links[0].Type, "GET")
	assert.EqualValues(t, response.Links[1].Rel, "ingredients")
	assert.EqualValues(t, response.Links[1].Href, "localhost:3000/pasteis/"+pastel.ID)
	assert.EqualValues(t, response.Links[1].Type, "DELETE")
}

func TestUpdate_Not_Found(t *testing.T) {
	pastelToUpdate := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(
		pastelToUpdate,
	)

	req, err := http.NewRequest("PUT", utils.BaseUri+"/"+pastelToUpdate.ID, &buf)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repository.NewPastelRepository(dataprovider.NewSqlite())).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusNotFound, resp.Code)

	var response model.ErrorResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.NotNil(t, response)
	assert.False(t, response.Success)
	assert.EqualValues(t, "registro não encontrado", response.Message)
	assert.Empty(t, response.Errors)
}

func TestUpdate_Validation_Error(t *testing.T) {
	pastelToUpdate := test.BuildPastelDomainWithIgredients("Pantaneiro", []string{"Carne"})

	pastelToUpdate.Name = ""
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(
		pastelToUpdate,
	)

	req, err := http.NewRequest("PUT", utils.BaseUri+"/"+pastelToUpdate.ID, &buf)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repository.NewPastelRepository(dataprovider.NewSqlite())).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusBadRequest, resp.Code)

	var response model.ErrorResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.NotNil(t, response)
	assert.False(t, response.Success)
	assert.EqualValues(t, "Validation error", response.Message)
	assert.NotNil(t, response.Errors)
	assert.EqualValues(t, "Name is required", response.Errors[0])
}
