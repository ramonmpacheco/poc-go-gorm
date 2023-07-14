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

func Init(repository dataprovider.IPastelRepository) {
	r := chi.NewRouter()
	createController := entrypoint.NewCreatePastelController(
		usecase.NewCreatePastelUseCase(repository),
	)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Route(utils.BaseUri, func(r chi.Router) {
		r.Post("/", createController.Create)
	})
	fmt.Println("About to start http server, port: 3000")
	http.ListenAndServe(":3000", r)
}
