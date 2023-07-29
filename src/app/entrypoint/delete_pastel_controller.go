package entrypoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
)

type deletePastelController struct {
	UseCase usecase.IDeletePastelUseCase
}

func NewDeletePastelController(useCase usecase.IDeletePastelUseCase) *deletePastelController {
	return &deletePastelController{
		UseCase: useCase,
	}
}

func (fpuc *deletePastelController) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := fpuc.UseCase.DeleteById(id)
	if err != nil {
		render.Status(r, model.GetStatusFrom(err))
		render.JSON(w, r, model.NewErrorResponse(err.Error(), nil))
		return
	}
	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, model.NewDeleteResponseSuccess(id))
}
