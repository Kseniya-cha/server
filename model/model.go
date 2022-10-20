package model

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

// общий конфиг
type Config struct {
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	ConfigDB `yaml:"database"`
}

// параметры логгера
type Logger struct {
	LogLevel string `yaml:"loglevel"`
}

// параметры сервера
type Server struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readtimeout"`  //time.Duration
	WriteTimeout time.Duration `yaml:"writetimeout"` //time.Duration
	IdleTimeout  time.Duration `yaml:"idletimeout"`  //time.Duration
}

// параметры базы данных
type ConfigDB struct {
	Port     string `yaml:"port" default:"5432"`
	Host     string `yaml:"host"`
	Db_name  string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// структура таблицы refresh_stream
// sql.Null* когда возможен null в столбце
type RefreshStreamWithNull struct {
	Id            int
	Auth          sql.NullString
	Ip            sql.NullString
	Stream        sql.NullString
	Run           sql.NullString
	Portsrv       string
	Sp            sql.NullString
	Camid         sql.NullString
	Record_status sql.NullBool
	Stream_status sql.NullBool
	Record_state  sql.NullBool
	Stream_state  sql.NullBool
}

type DBLog struct {
	Db  *sql.DB
	Log *logrus.Logger
}
