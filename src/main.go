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
	// carne := model.Ingredient{
	// 	Name: "Carne" + xid.New().String(),
	// 	Desc: "250 gramas",
	// }
	// queijo := model.Ingredient{
	// 	Name: "Queijo mussarela" + xid.New().String(),
	// 	Desc: "200 gramas",
	// }

	// ingredientes := []model.Ingredient{
	// 	carne, queijo,
	// }
	// carneQueijo := model.Pastel{
	// 	Name:        "Carne com queijo" + xid.New().String(),
	// 	Price:       9.50,
	// 	Ingredients: ingredientes,
	// }

	repository := repository.NewPastelRepository(dataprovider.NewDb())
	createHandler := entrypoint.Handler{UseCase: usecase.NewCreatePastelUseCase(repository)}
	findHandler := entrypoint.Handler{UseCase: usecase.NewFindPastelUseCase(repository)}
	r := chi.NewRouter()
	r.Post("/pastel", createHandler.Create)
	r.Get("/pastel", findHandler.Find)
	http.ListenAndServe(":3000", r)
}
