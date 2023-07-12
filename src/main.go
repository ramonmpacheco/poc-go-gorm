package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/postgres/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
)

func main() {
	r := chi.NewRouter()
	repository := repository.NewPastelRepository(dataprovider.NewDb())
	controller := entrypoint.NewCreatePastelController(usecase.NewCreatePastelUseCase(repository))
	r.Post("/pastel", controller.Create)
	http.ListenAndServe(":3000", r)
}
