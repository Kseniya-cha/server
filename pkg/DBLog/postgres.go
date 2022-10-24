package dblog

import (
	"database/sql"
	"fmt"

	"github.com/Kseniya-cha/server/constants"
	"github.com/Kseniya-cha/server/pkg/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// открывает базу данных и подключается к ней
func OpenDB(cfg config.Config, log *logrus.Logger) *sql.DB {
	sqlInfoOpen := fmt.Sprintf(constants.OpenDBConst,
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Db_name)

	db, err := sql.Open(constants.DriverName, sqlInfoOpen)
	LogBeforeDB(log, err, fmt.Sprintf(constants.OpenDBRespConst, cfg.Db_name))

	err = db.Ping()
	LogBeforeDB(log, err, fmt.Sprintf(constants.OpenDBPingRespConst, cfg.Db_name))

	return db
}

// закрывает базу данных
func CloseDB(db DBLog, cfg config.Config) {
	err := db.Db.Close()
	db.LogIFAction(err, fmt.Sprintf(constants.CloseDBRespConst, cfg.Db_name))
}
