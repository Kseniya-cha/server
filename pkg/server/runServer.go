package server

import (
	"fmt"
	"net/http"

	"github.com/Kseniya-cha/server/constants"
	dblog "github.com/Kseniya-cha/server/pkg/DBLog"
)

func RunServer(db dblog.DBLog, server *http.Server) {
	errLAS := server.ListenAndServe()
	if errLAS == http.ErrServerClosed {
		db.LogE(constants.ServCloseConst)
	} else if errLAS != nil {
		db.LogE(fmt.Sprintf(constants.ServErrConst, errLAS))
	}
}
