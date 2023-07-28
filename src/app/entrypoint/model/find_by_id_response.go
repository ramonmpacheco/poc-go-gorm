package model

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

type FindByIdResponse struct {
	Success bool           `json:"success"`
	Data    PastelResponse `json:"data"`
	Links   []link         `json:"_links,omitempty"`
}

func NewFindByIdSuccessResponse(pastel PastelResponse) FindByIdResponse {
	return FindByIdResponse{
		Success: true,
		Data:    pastel,
		Links:   getFindByIdResponseLinks(pastel.ID, utils.BaseUrl+utils.BaseUri),
	}
}

func getFindByIdResponseLinks(id, selfHref string) []link {
	links := make([]link, 0)
	links = append(links, link{Rel: "ingredients", Href: fmt.Sprintf("%s/%s", selfHref, id), Type: "DELETE"})
	return links
}
