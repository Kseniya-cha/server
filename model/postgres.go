package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// определяется уровень логирования
func (cfg Config) DefLogLevel() string {
	return cfg.LogLevel
}

// открывает базу данных и подключается к ней
func (cfg Config) OpenDB(log *logrus.Logger) *sql.DB {
	sqlInfoOpen := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Db_name)

	db, err := sql.Open("postgres", sqlInfoOpen)
	LogBeforeDB(log, err, fmt.Sprintf("open database %s", cfg.Db_name))

	err = db.Ping()
	LogBeforeDB(log, err, fmt.Sprintf("connection to database %s", cfg.Db_name))

	return db
}

// закрывает базу данных
func (cfg Config) CloseDB(db DBLog) {
	err := db.Db.Close()
	LogBeforeDB(db.Log, err, fmt.Sprintf("close database %s", cfg.Db_name))
}
