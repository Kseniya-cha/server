package server

import (
	"net/http"

	"github.com/Kseniya-cha/server/pkg/config"
)

// инициализация сервера с параметрами из конфига
func NewServer(cfg config.Config, router http.Handler) *http.Server {

	return &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
