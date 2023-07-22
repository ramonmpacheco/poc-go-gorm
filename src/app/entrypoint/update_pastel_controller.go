package entrypoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/converter"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/validator"
	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"
)

type updatePastelController struct {
	UseCase usecase.IUpdatePastelUseCase
}

func NewUpdatePastelController(useCase usecase.IUpdatePastelUseCase) *updatePastelController {
	return &updatePastelController{
		UseCase: useCase,
	}
}

func (upc *updatePastelController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var request model.UpdatePastelRequest
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewCreateResponse(false, "Erro ao processar, revise os dados enviados"))
		return
	}

	result, err := validator.ValidateStruct(request)
	if err != nil {
		render.Status(r, model.GetStatusFrom(err))
		render.JSON(w, r, model.NewErrorResponse("Validation error", result))
		return
	}

	err = upc.UseCase.Update(id, converter.ToPastelDomainFromUpdate(request))
	if err != nil {
		render.Status(r, model.GetStatusFrom(err))
		render.JSON(w, r, model.NewErrorResponse(err.Error(), nil))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.NewUpdateResponseSuccess(id))
}
