package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Kseniya-cha/server/constants"
	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/Kseniya-cha/server/pkg/logger"
	_ "github.com/lib/pq"
)

func CreateDBConnection(cfg *config.Config) *sql.DB {
	var dbcfg ConfigDB

	dbcfg.Port = cfg.Port
	dbcfg.Host = cfg.Host
	dbcfg.Db_name = cfg.Db_name
	dbcfg.User = cfg.User
	dbcfg.Password = cfg.Password

	dbcfg.Driver = cfg.Driver
	dbcfg.DBConnectionTimeoutSecond = 10 * time.Second //cfg.DBConnectionTimeoutSecond
	dbcfg.Log = logger.NewLog(cfg.LogLevel)

	return connectToDB(&dbcfg)
}

func connectToDB(dbcfg *ConfigDB) *sql.DB {
	var db *sql.DB

	// подключение к базе
	sqlInfo := fmt.Sprintf(constants.OpenDBConst,
		dbcfg.Host, dbcfg.Port, dbcfg.User, dbcfg.Password,
		dbcfg.Db_name)

	db, _ = sql.Open(dbcfg.Driver, sqlInfo)
	time.Sleep(time.Millisecond * 3)
	if db.Ping() == nil {
		return db
	}

	connLatency := time.Duration(10 * time.Millisecond)
	time.Sleep(connLatency * time.Millisecond)
	connTimeout := dbcfg.DBConnectionTimeoutSecond
	for t := connTimeout; t > 0; t-- {
		if db != nil {
			logger.LogDebug(dbcfg.Log, fmt.Sprintf("established opening of connection to DB at %v", time.Now()))
			return db
		}
		time.Sleep(time.Second * 3)
	}

	logger.LogError(dbcfg.Log, fmt.Sprintf("time waiting of DB connection exceeded limit: %v", connTimeout))
	return db
}

func CloseDBConnection(cfg *config.Config, db *sql.DB) {
	log := logger.NewLog(cfg.LogLevel)
	if err := db.Close(); err != nil {
		logger.LogError(log, fmt.Sprintf("Cannot close DB connection. Error: %v", err))
		return
	}
	logger.LogDebug(log, "established closing of connection to DB")
}
