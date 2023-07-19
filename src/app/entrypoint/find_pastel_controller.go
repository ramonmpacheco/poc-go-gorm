package entrypoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
)

type findPastelController struct {
	UseCase usecase.IFindPastelUseCase
}

func NewFindPastelController(useCase usecase.IFindPastelUseCase) *findPastelController {
	return &findPastelController{
		UseCase: useCase,
	}
}

func (fpuc *findPastelController) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fpuc.UseCase.FindById(id)
}
