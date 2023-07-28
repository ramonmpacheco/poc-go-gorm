package model

import (
	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
	Links   []link `json:"_links,omitempty"`
}

func NewDeleteResponseSuccess(id string) CreateResponse {
	return CreateResponse{
		Success: true,
		ID:      id,
		Message: "resource successfuly deleted",
		Links:   getDeleteResponseLinks(id, utils.BaseUrl+utils.BaseUri),
	}
}

func getDeleteResponseLinks(id, selfHref string) []link {
	links := make([]link, 0)
	return links
}
