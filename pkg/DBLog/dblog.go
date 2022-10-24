package dblog

import (
	"database/sql"

	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/sirupsen/logrus"
)

type DBLog struct {
	Db  *sql.DB
	Log *logrus.Logger
}

// инициализация нового объекта DBLog и функции закрытия файла out.log
func NewDBLog(cfg config.Config) (DBLog, func()) {
	log, close := NewLog(cfg.LogLevel)
	db := OpenDB(cfg, log)
	return DBLog{
		Db:  db,
		Log: log,
	}, close
}
