package entrypoint

import (
	"fmt"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create chamado")
}
