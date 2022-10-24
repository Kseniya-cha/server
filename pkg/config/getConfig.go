package config

import (
	"io"
	"os"

	"github.com/Kseniya-cha/server/constants"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var cfg *Config

func GetConfig() *Config {
	cfg = &Config{}

	// чтение конфига
	errRC := cleanenv.ReadConfig("config.yml", &cfg)

	// логгер с функцией отложенного закрытия файла логирования
	// log, closeLogFile := dblog.NewLog(cfg.LogLevel)
	// defer closeLogFile()

	file, err := os.OpenFile(constants.FileNameConst, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf(constants.OpenFileErrConst, err)
	}
	defer file.Close()

	log := &logrus.Logger{
		Out:   io.MultiWriter(file, os.Stdout),
		Level: initLogLevel(cfg.LogLevel),
		Formatter: &easy.Formatter{
			TimestampFormat: constants.ServTimestampFormatConst,
			LogFormat:       constants.ServLogFormatConst,
		},
	}

	// проверка ошибки чтения конфига
	if errRC != nil {
		log.Fatal(constants.ReadConfigConst)
	}

	// возвращает объект *Config
	return cfg
}

func initLogLevel(level string) logrus.Level {
	switch level {
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
