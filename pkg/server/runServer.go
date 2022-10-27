package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/internal/refreshStream/controller"
	"github.com/Kseniya-cha/server/internal/refreshStream/repository"
	"github.com/Kseniya-cha/server/internal/refreshStream/usecase"
	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/Kseniya-cha/server/pkg/database"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type app struct {
	http                 *http.Server
	cfg                  *config.Config
	log                  *logrus.Logger
	db                   *sql.DB
	refreshStreamUseCase refreshStream.RefreshStreamUseCase
}

func NewApp(cfg *config.Config) *app {
	log := logger.NewLog(cfg.LogLevel)
	db := database.CreateDBConnection(cfg)
	repo := repository.NewRefreshStreamRepository(db, log)

	return &app{
		cfg:                  cfg,
		db:                   db,
		log:                  log,
		refreshStreamUseCase: usecase.NewRefreshStreamUseCase(repo, db, log),
	}
}

func (a *app) RunServer() error {

	router := mux.NewRouter()

	// регистрация маршрутов
	controller.RegisterRouter(router, a.refreshStreamUseCase, a.db, a.log)

	// инициализация сервера
	a.http = &http.Server{
		Addr:         a.cfg.Addr,
		Handler:      router,
		ReadTimeout:  a.cfg.ReadTimeout,
		WriteTimeout: a.cfg.WriteTimeout,
		IdleTimeout:  a.cfg.IdleTimeout,
	}

	// запуск сервера
	err := a.http.ListenAndServe()
	if err == http.ErrServerClosed {
		logger.LogError(a.log, ServCloseConst)
	} else if err != nil {
		logger.LogError(a.log, fmt.Sprintf(ServErrConst, err))
	}

	// завершение работы сервера
	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err := a.http.Shutdown(ctx); err != nil {
		return fmt.Errorf(ShutdownErrConst, err)
	}

	logger.LogInfo(a.log, ShutdownOkConst)

	return nil
}

// ожидание прерывающего сигнала
func (a *app) GracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	logger.LogFatal(a.log, fmt.Sprintf(SigConst, sig))
}
