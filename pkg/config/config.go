package config

import (
	"time"
)

// общий конфиг
type Config struct {
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	ConfigDB `yaml:"database"`
	TestConn `yaml:"testconn"`
}

// параметры логгера
type Logger struct {
	LogLevel string `yaml:"loglevel"`
}

// параметры сервера
type Server struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readtimeout"`
	WriteTimeout time.Duration `yaml:"writetimeout"`
	IdleTimeout  time.Duration `yaml:"idletimeout"`
}

// параметры базы данных
type ConfigDB struct {
	Port     string `yaml:"port" default:"5432"`
	Host     string `yaml:"host"`
	Db_name  string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// промежуток времени, через который происходит
// проверка соединения с базой
type TestConn struct {
	TimeTest time.Duration `yaml:"timetest"`
}

type ConfigInt interface {
	Config
}
