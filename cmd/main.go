package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/EclipIIse/exchanger/internal/service"
	transport "github.com/EclipIIse/exchanger/internal/transport/http"
	"github.com/EclipIIse/exchanger/internal/transport/http/handlers"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	srv := service.New(logger)
	handler := handlers.New(logger, srv)
	server := transport.NewServer("127.0.0.1:8000").WithHandler(handler)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	go func() {
		if err := server.Run(); err != nil {
			logger.Fatal().Err(err).Msg("server starting error")
		}
	}()

	<-shutdown
}
