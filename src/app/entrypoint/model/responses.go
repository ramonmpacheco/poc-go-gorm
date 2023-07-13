package model

import "fmt"

type createResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
	Links   []link `json:"_links,omitempty"`
}

type link struct {
	Rel  string `json:"relationship"`
	Href string `json:"href"`
	Type string `json:"type"`
}

func NewCreateResponseSuccess(id, selfHref string) createResponse {
	return createResponse{
		Success: true,
		ID:      id,
		Message: "resource successfuly created",
		Links:   getLinks(id, selfHref),
	}
}

func NewCreateResponse(success bool, message string) createResponse {
	return createResponse{
		Success: success,
		Message: message,
	}
}

func getLinks(id, selfHref string) []link {
	links := make([]link, 0)
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "GET"})
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "DELETE"})
	return links
}

type errorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func NewErrorResponse(message string, errs []string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
		Errors:  errs,
	}
}
