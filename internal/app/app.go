package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
	"github.com/romankravchuk/currency-exchanger/internal/config"
	"github.com/romankravchuk/currency-exchanger/internal/handler"
	"github.com/romankravchuk/pastebin/pkg/httpserver"
	"github.com/romankravchuk/pastebin/pkg/log"
)

func Run(cfg *config.Config) {
	logger := log.New(os.Stdout, log.Stol(cfg.Logger.Level))

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(handler.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(handler.MethodNotAllowed)

	handler.MountCurrencyRouter(router, logger)

	server := httpserver.New(router)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	logger.Info("starting HTTP server", log.FF{
		{Key: "port", Value: cfg.Server.Port},
	})

	select {
	case s := <-interrupt:
		logger.Info("catch interrupt signal", log.FF{
			{Key: "signal", Value: s.String()},
		})
	case err := <-server.Notify():
		logger.Error("HTTP server error", err, nil)
	}

	err := server.Shutdown()
	if err != nil {
		logger.Fatal("HTTP server shutdown error", err, nil)
	}
}
