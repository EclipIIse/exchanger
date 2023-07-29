package http

import (
	"net/http"

	"github.com/EclipIIse/exchanger/internal/transport/http/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/currency", h.GetCurrency).Methods(http.MethodGet)

	return r
}
