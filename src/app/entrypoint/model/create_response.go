package model

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

type CreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
	Links   []link `json:"_links,omitempty"`
}

func NewCreateResponseSuccess(id string) CreateResponse {
	return CreateResponse{
		Success: true,
		ID:      id,
		Message: "resource successfuly created",
		Links:   getCreateResponseLinks(id, utils.BaseUrl+utils.BaseUri),
	}
}

func NewCreateResponse(success bool, message string) CreateResponse {
	return CreateResponse{
		Success: success,
		Message: message,
	}
}

func getCreateResponseLinks(id, selfHref string) []link {
	links := make([]link, 0)
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "GET"})
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "DELETE"})
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "PUT"})
	return links
}
