package entrypoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/converter"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
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
	pastel, err := fpuc.UseCase.FindById(id)
	if err != nil {
		render.Status(r, model.GetStatusFrom(err))
		render.JSON(w, r, model.NewCreateResponse(false, err.Error()))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.NewFindByIdSuccessResponse(converter.ToPastelResponse(*pastel)))
}
