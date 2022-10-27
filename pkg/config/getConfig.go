package config

import (
	"fmt"

	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg *Config

func GetConfig() *Config {
	cfg = &Config{}

	// чтение конфига
	errRC := cleanenv.ReadConfig(ConfigFileNameConst, cfg)

	// обработка ошибки чтения конфига
	log := logger.NewLog(cfg.LogLevel)
	if errRC != nil {
		logger.LogError(log, fmt.Sprintf(ReadConfigErrConst, errRC))
	} else {
		logger.LogDebug(log, ReadConfigEOkConst)
	}

	return cfg
}
