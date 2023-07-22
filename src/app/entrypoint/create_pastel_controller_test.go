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

func TestCreate_success(t *testing.T) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(
		test.BuildCreatePastelRequest("Boi ralado", []string{"Carne"}),
	)

	req, err := http.NewRequest("POST", utils.BaseUri, &buf)
	assert.Nil(t, err)
	resp := httptest.NewRecorder()

	rest.InitRoutes(repository.NewPastelRepository(dataprovider.NewSqlite())).
		ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusCreated, resp.Code)

	var response model.CreateResponse
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.NotNil(t, response.ID)
	assert.True(t, response.Success)
	assert.EqualValues(t, response.Message, "resource successfuly created")
	assert.EqualValues(t, response.Links[0].Rel, "ingredients")
	assert.EqualValues(t, response.Links[0].Href, "localhost:3000/pasteis/"+response.ID)
	assert.EqualValues(t, response.Links[0].Type, "GET")

}

func TestCreate_validation_error(t *testing.T) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(
		test.BuildCreatePastelRequest("", []string{"Carne"}),
	)

	req, err := http.NewRequest("POST", utils.BaseUri, &buf)
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
	assert.EqualValues(t, "Name is required", response.Errors[0])
}
