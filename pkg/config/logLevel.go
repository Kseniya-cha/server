package config

// определяется уровень логирования
func (cfg Config) DefLogLevel() string {
	return cfg.LogLevel
}
