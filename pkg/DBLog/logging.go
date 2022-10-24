package dblog

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// проверка ошибки в случае, когда структура DBLog не инициализирована
func LogBeforeDB(log *logrus.Logger, err error, action string) {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Infof("%s - ok", action)
	}
}

// проверка ошибки: если есть - фатал, если нет: "действие action - ок"
func (db DBLog) LogIFAction(err error, action string) {
	if err != nil {
		db.Log.Fatal(err)
	} else {
		db.Log.Infof("%s - ok", action)
	}
}

// проверка ошибки err: если есть - фатал
func (db DBLog) LogPrintFat(err error) {
	if err != nil {
		db.Log.Fatal(err)
	}
}

// проверка ошибки: если есть - печать сообщения msg в консоль и
// в терминал, затем фатал (err), если нет: "success действие msg"
func (db DBLog) LogWriteIF(w http.ResponseWriter, err error, msg string) {
	if err != nil {
		db.LogWrite(w, fmt.Sprintf("fail %s", msg))
		db.Log.Fatal(err)
	} else {
		db.LogWrite(w, fmt.Sprintf("success %s", msg))
	}
}

// выводит сообщение msg с флажком "Fatal"
func (db DBLog) LogF(msg string) {
	db.Log.Fatalf(msg)
}

// выводит сообщение msg с флажком "Error"
func (db DBLog) LogE(msg string) {
	db.Log.Errorf(msg)
}

// выводит сообщение msg с флажком "Info"
func (db DBLog) LogI(msg string) {
	db.Log.Infof(msg)
}

// печать сообщения msg в консоль и в терминал
func (db DBLog) LogWrite(w http.ResponseWriter, msg string) {
	db.LogI(msg)
	io.WriteString(w, fmt.Sprintf("%s\n", msg))
}
