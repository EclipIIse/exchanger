package handlers

import (
	"encoding/json"
	"github.com/EclipIIse/exchanger/internal/models"
	"net/http"

	"github.com/rs/zerolog"
)

type Service interface {
	GetLocalCurrencyHistory() error
	GetCurrency() (*models.CurrencyResponse, error)
}

type Handler struct {
	log     zerolog.Logger
	service Service
}

func New(log zerolog.Logger, service Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := h.service.GetCurrency()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		h.log.Error().Err(err).Msg("filed to marshal response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
