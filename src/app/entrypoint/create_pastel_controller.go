package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"

	"github.com/go-chi/render"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request model.CreatePastelRequest
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		fmt.Println(err.Error())
	}

}
