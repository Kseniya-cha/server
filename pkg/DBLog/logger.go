package dblog

import (
	"io"
	"os"

	"github.com/Kseniya-cha/server/constants"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type Logger struct {
	*logrus.Logger
}

func NewLog(level string) (*logrus.Logger, func()) {
	file, err := os.OpenFile(constants.FileNameConst, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf(constants.OpenFileErrConst, err)
	}
	close := func() {
		file.Close()
	}

	log := &logrus.Logger{
		Out:   io.MultiWriter(file, os.Stdout),
		Level: initLogLevel(level),
		Formatter: &easy.Formatter{
			TimestampFormat: constants.ServTimestampFormatConst,
			LogFormat:       constants.ServLogFormatConst,
		},
	}

	return log, close
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
