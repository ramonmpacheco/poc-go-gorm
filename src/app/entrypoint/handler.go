package entrypoint

import "github.com/ramonmpacheco/poc-go-gorm/domain/usecase"

type Handler struct {
	UseCase usecase.PastelUseCase
}
