package methods

import (
	"fmt"
	"net/http"

	"github.com/Kseniya-cha/server/logger"
	"github.com/sirupsen/logrus"
)

// запуск сервера и обработка ошибок
func RunServer(log *logrus.Logger, server *http.Server) {
	errLAS := server.ListenAndServe()
	if errLAS == http.ErrServerClosed {
		logger.LogE(log, "server closed")
	} else if errLAS != nil {
		logger.LogE(log, fmt.Sprintf("error listening for server: %s", errLAS))
	}
}
