package storage

import "github.com/rs/zerolog"

type Storage interface {
	GetLocalCurrencyHistory() error
}

type storage struct {
	log zerolog.Logger
}

func (s *storage) GetLocalCurrencyHistory() error {
	return nil
}
