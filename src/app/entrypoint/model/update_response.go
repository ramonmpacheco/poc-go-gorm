package model

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

type UpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Links   []link `json:"_links,omitempty"`
}

func NewUpdateResponseSuccess(id string) UpdateResponse {
	return UpdateResponse{
		Success: true,
		Message: "resource successfuly updated",
		Links:   getUpdateResponseLinks(id, utils.BaseUrl+utils.BaseUri),
	}
}

func getUpdateResponseLinks(id, selfHref string) []link {
	links := make([]link, 0)
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "GET"})
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "DELETE"})
	return links
}
