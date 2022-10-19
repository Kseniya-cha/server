package logger

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func LogIFAction(log *logrus.Logger, err error, action string) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	} else {
		log.Infof("%s - ok", action)
	}
}

func LogE(log *logrus.Logger, msg string) {
	log.Errorf(msg)
}

func LogI(log *logrus.Logger, msg string) {
	log.Infof(msg)
}

func LogF(log *logrus.Logger, msg string) {
	log.Fatalf(msg)
}

func LogWrite(log *logrus.Logger, w http.ResponseWriter, msg string) {
	LogI(log, msg)
	io.WriteString(w, fmt.Sprintf("%s\n", msg))
}
