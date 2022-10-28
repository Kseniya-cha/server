package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/Kseniya-cha/server/pkg/logger"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func CreateDBConnection(cfg *config.Config) *sql.DB {
	var dbcfg ConfigDB

	dbcfg.Port = cfg.Port
	dbcfg.Host = cfg.Host
	dbcfg.Db_name = cfg.Db_name
	dbcfg.User = cfg.User
	dbcfg.Password = cfg.Password

	dbcfg.Driver = cfg.Driver
	dbcfg.DBConnectionTimeoutSecond = cfg.DBConnectionTimeoutSecond
	dbcfg.Log = logger.NewLog(cfg.LogLevel)

	return connectToDB(&dbcfg)
}

func connectToDB(dbcfg *ConfigDB) *sql.DB {
	var dbSQL *sql.DB
	var dbGORM *gorm.DB

	// подключение к базе
	sqlInfo := fmt.Sprintf(DBInfoConst,
		dbcfg.Host, dbcfg.Port, dbcfg.User, dbcfg.Password,
		dbcfg.Db_name)

	dbSQL, err := sql.Open(dbcfg.Driver, sqlInfo)
	if err != nil {
		logger.LogError(dbcfg.Log, fmt.Sprintf(OpenDBErrConst, "sql"))
	}

	// dbGORM, err = gorm.Open(postgres.New(postgres.Config{
	// 	Conn: dbSQL,
	// }), &gorm.Config{})
	// if err != nil {
	// 	logger.LogError(dbcfg.Log, fmt.Sprintf(OpenDBErrConst, "gorm"))
	// }

	// автомиграция
	// dbGORM.AutoMigrate(&refreshStream.RefreshStreamWithNull{})

	// проверка подключения
	time.Sleep(time.Millisecond * 3)
	if dbGORM == nil {
		logger.LogDebug(dbcfg.Log, fmt.Sprintf(ConnectToDBOkConst, dbcfg.Db_name))
		return dbSQL
	} else {
		logger.LogError(dbcfg.Log, ConnectToDBErrConst)
	}

	connLatency := time.Duration(10 * time.Millisecond)
	time.Sleep(connLatency * time.Millisecond)
	connTimeout := dbcfg.DBConnectionTimeoutSecond
	for t := connTimeout; t > 0; t-- {
		if dbGORM != nil {
			return dbSQL
		}
		time.Sleep(time.Second * 3)
	}

	logger.LogError(dbcfg.Log, fmt.Sprintf(WaitForBDErrConst, connTimeout))
	return dbSQL
}

// как закрыть? или не надо закрывать, раз нет такой функции?
func CloseDBConnection(cfg *config.Config, dbSQL *sql.DB) {
	log := logger.NewLog(cfg.LogLevel)
	if err := dbSQL.Close(); err != nil {
		logger.LogError(log, fmt.Sprintf(CloseDBErrConst, err))
		return
	}
	logger.LogDebug(log, CloseDBOkConst)
}
