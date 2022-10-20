package model

func NewDBLog(cfg Config, level string) (DBLog, func()) {
	log, close := DefLog(level)
	db := cfg.OpenDB(log)
	return DBLog{
		Db:  db,
		Log: log,
	}, close
}
