package database

import (
	"time"

	"github.com/sirupsen/logrus"
)

type ConfigDB struct {
	Port     string
	Host     string
	Db_name  string
	User     string
	Password string

	Driver                    string
	DBConnectionTimeoutSecond time.Duration
	// DBConnectionLatencyMillisecond time.Duration
	Log *logrus.Logger
}
