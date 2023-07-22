package entrypoint

import (
	"net/http"

	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/converter"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/model"
	"github.com/ramonmpacheco/poc-go-gorm/app/entrypoint/validator"

	"github.com/ramonmpacheco/poc-go-gorm/domain/usecase"

	"github.com/go-chi/render"
)

type createPastelController struct {
	UseCase usecase.ICreatePastelUseCase
}

func NewCreatePastelController(useCase usecase.ICreatePastelUseCase) *createPastelController {
	return &createPastelController{
		UseCase: useCase,
	}
}

func (cpc *createPastelController) Create(w http.ResponseWriter, r *http.Request) {
	var request model.CreatePastelRequest
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

	id, err := cpc.UseCase.Create(converter.ToPastelDomainFromCreate(request))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewCreateResponse(false, err.Error()))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, model.NewCreateResponseSuccess(id))
}
