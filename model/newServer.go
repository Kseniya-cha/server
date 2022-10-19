package model

import (
	"net/http"
)

// инициализация сервера с параметрами из конфига
func (cfg Config) NewServer(router http.Handler) *http.Server {

	return &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
