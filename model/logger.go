package model

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func DefLog(level string) (*logrus.Logger, func()) {
	file, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	close := func() {
		file.Close()
	}

	log := &logrus.Logger{
		Out:   io.MultiWriter(file, os.Stdout),
		Level: initLogLevel(level),
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: (%time%) %msg%\n",
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
