package logger

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// выводит сообщение msg на уровне "Fatal"
func LogFatal(log *logrus.Logger, msg interface{}) {
	log.Fatalf(fmt.Sprintf("%v", msg))
}

// выводит сообщение msg на уровне "Error"
func LogError(log *logrus.Logger, msg interface{}) {
	log.Errorf(fmt.Sprintf("%v", msg))
}

// выводит сообщение msg на уровне "Warn"
func LogWarn(log *logrus.Logger, msg interface{}) {
	log.Warnf(fmt.Sprintf("%v", msg))
}

// выводит сообщение msg на уровне "Info"
func LogInfo(log *logrus.Logger, msg interface{}) {
	log.Infof(fmt.Sprintf("%v", msg))
}

// выводит сообщение msg на уровне "Debug"
func LogDebug(log *logrus.Logger, msg interface{}) {
	log.Debugf(fmt.Sprintf("%v", msg))
}

// выводит сообщение msg на уровне "Trace"
func LogTrace(log *logrus.Logger, msg interface{}) {
	log.Tracef(fmt.Sprintf("%v", msg))
}

//========================================================================================

// проверка ошибки: если есть - фатал, если нет: "действие action - ок"
func LogFatalOrAction(log *logrus.Logger, err error, action string) {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Infof("%s - ok", action)
	}
}

// проверка ошибки err: если есть - фатал
func LogPrintFatal(log *logrus.Logger, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//========================================================================================

// печать сообщения msg в консоль и в терминал
func LogWriteInfo(log *logrus.Logger, w http.ResponseWriter, msg string) {
	LogInfo(log, msg)
	io.WriteString(w, fmt.Sprintf("%v", msg))
}

// печать сообщения msg в консоль и в терминал
func LogWriteDebug(log *logrus.Logger, w http.ResponseWriter, msg string) {
	LogDebug(log, msg)
	io.WriteString(w, fmt.Sprintf("%v", msg))
}

// проверка ошибки: если есть - печать сообщения msg в консоль и
// в терминал, затем фатал (err), если нет: "success действие msg"
func LogWrite(log *logrus.Logger, w http.ResponseWriter, err error, msg string) {
	if err != nil {
		LogWriteInfo(log, w, fmt.Sprintf("fail %s", msg))
		log.Fatal(err)
	} else {
		LogWriteInfo(log, w, fmt.Sprintf("success %s", msg))
	}
}
