package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint"
	"github.com/ramonmpacheco/poc-go-gorm/domain/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
	"github.com/ramonmpacheco/poc-go-gorm/utils"
)

func InitRoutes(repository dataprovider.IPastelRepository) *chi.Mux {
	r := chi.NewRouter()
	createController := entrypoint.NewCreatePastelController(
		usecase.NewCreatePastelUseCase(repository),
	)
	findController := entrypoint.NewFindPastelController(
		usecase.NewFindPastelUseCase(repository),
	)
	updateController := entrypoint.NewUpdatePastelController(
		usecase.NewUpdatePastelUseCase(repository),
	)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Route(utils.BaseUri, func(r chi.Router) {
		r.Post("/", createController.Create)
		r.Get("/{id}", findController.FindById)
		r.Put("/{id}", updateController.Update)
	})
	return r
}

func Start(r *chi.Mux) {
	fmt.Println("About to start http server, port: 3000")
	http.ListenAndServe(":3000", r)
}
