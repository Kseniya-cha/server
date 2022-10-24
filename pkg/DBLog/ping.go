package dblog

import (
	"time"

	"github.com/Kseniya-cha/server/constants"
	"github.com/Kseniya-cha/server/pkg/config"
)

func PingDB(db DBLog, cfg config.Config) {
	for {
		time.Sleep(cfg.TestConn.TimeTest)
		err := db.Db.Ping()
		if err != nil {
			db.LogE(constants.ConnErrConst)
		} else {
			db.LogI(constants.ConnOkConst)
		}
	}
}
