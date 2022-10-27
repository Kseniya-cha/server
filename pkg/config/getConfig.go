package config

import (
	"fmt"
	"os"

	"github.com/Kseniya-cha/server/constants"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg *Config

func GetConfig() *Config {
	cfg = &Config{}

	// чтение конфига
	errRC := cleanenv.ReadConfig("config.yml", cfg)

	file, err := os.OpenFile(constants.FileNameConst, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log := logger.NewLog(cfg.LogLevel)
	if err != nil {
		logger.LogFatal(log, fmt.Sprintf(constants.OpenFileErrConst, err))
	} else {
		logger.LogDebug(log, "success open file out.log")
	}
	defer file.Close()

	if errRC != nil {
		logger.LogFatal(log, fmt.Sprintf(constants.ReadConfigConst, errRC))
	} else {
		logger.LogDebug(log, "success read config.yml")
	}

	return cfg
}
