package methods

import (
	"fmt"
	"net/http"

	"github.com/Kseniya-cha/server/model"
)

// запуск сервера и обработка ошибок
func RunServer(db model.DBLog, server *http.Server) {
	errLAS := server.ListenAndServe()
	if errLAS == http.ErrServerClosed {
		db.LogE("server closed")
	} else if errLAS != nil {
		db.LogE(fmt.Sprintf("error listening for server: %s", errLAS))
	}
}
