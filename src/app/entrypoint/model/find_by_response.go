package model

import (
	"fmt"
	"time"

	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

type FindByIdResponse struct {
	Success bool           `json:"success"`
	Data    PastelResponse `json:"data"`
	Links   []link         `json:"_links,omitempty"`
}

type PastelResponse struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Price       float32              `json:"price"`
	Ingredients []IngredientResponse `json:"ingredients"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type IngredientResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
