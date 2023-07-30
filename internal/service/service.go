package service

import (
	"github.com/EclipIIse/exchanger/internal/models"
	"github.com/rs/zerolog"
)

type Storage interface {
	GetLocalCurrencyHistory() error
}

type ExchangeClient interface {
	GetCurrency() (*models.CurrencyResponse, error)
}

type Service struct {
	log    zerolog.Logger
	client ExchangeClient
}

func (s *Service) GetCurrency() (*models.CurrencyResponse, error) {
	data, err := s.client.GetCurrency()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) GetLocalCurrencyHistory() error {
	return nil
}

func New(log zerolog.Logger, client ExchangeClient) *Service {
	return &Service{
		log:    log,
		client: client,
	}
}
