package dblog

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Kseniya-cha/server/constants"
)

// структура или интерфейс с дблогом,
// чтобы не было зависимости от других пакетов(?)
func GracefulShutdown(db DBLog) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	db.Log.Fatalf(constants.SigConst, sig)
}
