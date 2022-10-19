package methods

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Kseniya-cha/server/logger"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// 'http://localhost:3333/
func RootHF(log *logrus.Logger) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.LogI(log, "get page \"/\"")
		fmt.Printf("this is root page in terminal\n")
		io.WriteString(w, "this is root page in consol\n")
	}
}

func SmthHF() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("работает")
		fmt.Println(mux.Vars(r))
	}
}

// 'http://localhost:3333/api/get/'
func GetAllHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger, show bool) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if show {
			for _, str := range SelectContextRS(ctx, db, log) {
				log.Infof("%#v\n", str)
				io.WriteString(w, fmt.Sprintf("%#v\n", str))
			}
		}
		logger.LogWrite(log, w, "success select all rows!")
	}
}

// показать все строки по значению id
// curl 'http://localhost:3333/api/get/3/'
func GetIdHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger, show bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := mux.Vars(r)["ID"]
		val, errVal := strconv.Atoi(value)
		logger.LogIFAction(log, errVal, "converse value to int")

		if show {
			for _, str := range GetIdContextRS(ctx, db, log, val) {
				log.Infof("%#v\n", str)
				io.WriteString(w, fmt.Sprintf("%#v\n", str))
			}
		}
		logger.LogWrite(log, w, "success select!")
	}
}

// 'http://localhost:3333/api/delete/3/'
func DeleteIdHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		value := mux.Vars(r)["ID"]
		val, errVal := strconv.Atoi(value)
		logger.LogIFAction(log, errVal, "converse value to int")

		DeleteContext(ctx, db, log, "id", val)
		logger.LogWrite(log, w, "success delete!")
	}
}

// вставить строку
// http://localhost:3333/patch?id=3&ip=aaa&stream=bbb&sp=ccc
func PostHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// извлечение значений из строки запроса -> map[string]string
		colVals := mux.Vars(r)
		// преобразование id к типу int
		idstr := colVals["ID"]
		id, errVal := strconv.Atoi(idstr)
		logger.LogIFAction(log, errVal, "converse id to int")

		InsertContext(ctx, db, log, `"id"`, fmt.Sprintf(`'%d'`, id))
	}
}

// частичное изменение строки по id
func PatchHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// UpdateContext()
	}
}

// изменение всех колонок строки по id
// http://localhost:3333/api/put/6/auth/ip/stream/run/port/sp/cam/true/false/false/true/
func PutHF(ctx context.Context, db *sql.DB,
	log *logrus.Logger) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// извлечение значений из строки запроса -> map[string]string
		colVals := mux.Vars(r)
		// преобразование id к типу int
		idstr := colVals["ID"]
		id, errVal := strconv.Atoi(idstr)
		logger.LogIFAction(log, errVal, "converse id to int")

		for col, val := range colVals {
			// названия колонок преобразуются к нижнему регистру
			col = strings.ToLower(col)
			// пропускается колонка с id
			if col == "id" {
				continue
			} else if col == "record_status" {
				logParseUpdateBool(id, col, val, colVals, ctx, db, log, w)
			} else if col == "stream_status" {
				logParseUpdateBool(id, col, val, colVals, ctx, db, log, w)
			} else if col == "record_state" {
				logParseUpdateBool(id, col, val, colVals, ctx, db, log, w)
			} else if col == "stream_state" {
				logParseUpdateBool(id, col, val, colVals, ctx, db, log, w)
			} else {
				UpdateContext(ctx, db, log, col, val, id)
				logger.LogWrite(log, w, fmt.Sprintf("success put: column %v - value %v", col, val))
			}
		}
	}
}

// преобразование строки к Bool, проверка ошибки, выполнение put-запроса
// и сообщение об успешности выполнения функции
func logParseUpdateBool(id int, col, val string, colVals map[string]string,
	ctx context.Context, db *sql.DB, log *logrus.Logger, w http.ResponseWriter) {

	sst, err := strconv.ParseBool(colVals[strings.ToUpper(col)])
	if err != nil {
		logger.LogE(log, "can not converse to bool")
	}
	UpdateContext(ctx, db, log, col, sst, id)
	logger.LogWrite(log, w, fmt.Sprintf("success put: column %v - value %v", col, val))
}
