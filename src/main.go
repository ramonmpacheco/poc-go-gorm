package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/postgres/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

func main() {
	r := chi.NewRouter()
	repository := repository.NewPastelRepository(dataprovider.NewDb())
	controller := entrypoint.NewCreatePastelController(usecase.NewCreatePastelUseCase(repository))

	r.Use(middleware.RequestID)
	r.Route(utils.BaseUri, func(r chi.Router) {
		r.Post("/", controller.Create)
	})

	http.ListenAndServe(":3000", r)
}
