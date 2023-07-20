package model

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
