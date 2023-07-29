package service

import "github.com/rs/zerolog"

type Storage interface {
	GetLocalCurrencyHistory() error
}

type Service struct {
	log zerolog.Logger
}

func (s *Service) GetCurrency() (string, error) {
	return "Implement me", nil
}

func (s *Service) GetLocalCurrencyHistory() error {
	return nil
}

func New(log zerolog.Logger) *Service {
	return &Service{
		log: log,
	}
}
