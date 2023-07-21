package model

import (
	"net/http"
)

type ErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func NewErrorResponse(message string, errs []string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errs,
	}
}

func GetStatusFrom(err error) int {
	status := map[string]int{
		"invalid data":            http.StatusBadRequest,
		"registro não encontrado": http.StatusNotFound,
		"registro já existente, verifique os dados enviados": http.StatusBadRequest,
	}
	code, wasFound := status[err.Error()]
	if !wasFound {
		return http.StatusInternalServerError
	}
	return code
}
