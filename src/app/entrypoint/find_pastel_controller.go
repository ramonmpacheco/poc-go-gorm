package entrypoint

import (
	"fmt"
	"net/http"
)

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Find chamado")
}
