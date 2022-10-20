package methods

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Kseniya-cha/server/model"
	"github.com/gorilla/mux"
)

// http://localhost:3333/
func RootHF(db model.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db.LogI("get page \"/\"")
		fmt.Printf("this is root page in terminal\n")
		io.WriteString(w, "this is root page in consol\n")
	}
}

// http://localhost:3333/api/get/
func GetAllHF(ctx context.Context, db model.DBLog,
	show bool) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		jData, err := SelectContextRS(ctx, db)
		db.LogPrintFat(err)

		if show {
			db.LogWrite(w, fmt.Sprintf(string(jData)))
		}
		db.LogWrite(w, "success select all rows!")
	}
}

// показать все строки по значению id
// http://localhost:3333/api/get/3/
func GetIdHF(ctx context.Context, db model.DBLog,
	show bool) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		value := mux.Vars(r)["ID"]
		val, err := strconv.Atoi(value)
		db.LogPrintFat(err)

		jData, err := GetIdContextRS(ctx, db, val)
		db.LogPrintFat(err)

		if show {
			db.LogWrite(w, fmt.Sprintf(string(jData)))
		}
	}
}

// http://localhost:3333/api/delete/3/
func DeleteIdHF(ctx context.Context, db model.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		value := mux.Vars(r)["ID"]
		val, err := strconv.Atoi(value)
		db.LogPrintFat(err)

		err = DeleteContext(ctx, db, val)
		db.LogWriteIF(w, err, "delete")
	}
}

// вставить строку
// http://localhost:3333/api/post/auth/ip/stream/run/port/sp/cam/true/false/false/true/
func PostHF(ctx context.Context, db model.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// извлечение значений из строки запроса -> map[string]string
		colVals := mux.Vars(r)

		allcols := `"auth","ip","stream","run","portsrv","sp","camid","record_status","stream_status","record_state","stream_state"`
		allvalues := fmt.Sprintf(`'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'`,
			colVals["AUTH"], colVals["IP"], colVals["STREAM"], colVals["RUN"], colVals["PORTSRV"],
			colVals["SP"], colVals["CAMID"], colVals["RECORD_STATUS"], colVals["STREAM_STATUS"],
			colVals["RECORD_STATE"], colVals["STREAM_STATE"])
		err := InsertContext(ctx, db, allcols, allvalues)
		db.LogWriteIF(w, err, "insert")
	}
}

// изменение всех колонок строки по id
// http://localhost:3333/api/put/6/auth/ip/stream/run/port/sp/cam/true/false/false/true/
func PutHF(ctx context.Context, db model.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// извлечение значений из строки запроса -> map[string]string
		colVals := mux.Vars(r)

		// преобразование id к типу int
		idstr := colVals["ID"]
		id, err := strconv.Atoi(idstr)
		db.LogIFAction(err, "converse id to int")

		for col, val := range colVals {
			// названия колонок преобразуются к нижнему регистру
			col = strings.ToLower(col)

			// пропускается колонка с id
			if col == "id" {
				continue
				// для колонок, значение которых должно быть Bool,
				// сначала производится преобразование извлечённых
				// из строки запроса данных
			} else if col == "record_status" {
				err := logParseUpdateBoolUp(id, col, val, colVals, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "stream_status" {
				err := logParseUpdateBoolUp(id, col, val, colVals, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "record_state" {
				err := logParseUpdateBoolUp(id, col, val, colVals, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "stream_state" {
				err := logParseUpdateBoolUp(id, col, val, colVals, ctx, db, w)
				db.LogPrintFat(err)
			} else {
				err := UpdateContext(ctx, db, col, val, id)
				db.LogPrintFat(err)
			}
		}
	}
}

// частичное изменение строки по id
// http://localhost:3333/api/patch?id=125&auth=auth1&ip=10.2&stream=&run=&portsvr=poooort&sp=&camid=&record_status=true&stream_status=false&record_state=false&stream_state=true
func PatchHF(ctx context.Context, db model.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		row := scanRow(r)
		id, err := strconv.Atoi(row["id"])
		db.LogIFAction(err, "converse id to int")

		for col, val := range row {
			// пропускается колонка с id
			if col == "id" {
				continue
			} else if col == "record_status" && val != "" {
				err := logParseUpdateBoolLow(id, col, val, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "stream_status" && val != "" {
				err := logParseUpdateBoolLow(id, col, val, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "record_state" && val != "" {
				err := logParseUpdateBoolLow(id, col, val, ctx, db, w)
				db.LogPrintFat(err)
			} else if col == "stream_state" && val != "" {
				err := logParseUpdateBoolLow(id, col, val, ctx, db, w)
				db.LogPrintFat(err)
			} else {
				err := UpdateContext(ctx, db, col, val, id)
				db.LogPrintFat(err)
			}
		}
	}
}

// преобразование строки к Bool, проверка ошибки, выполнение put-запроса
// и сообщение об успешности выполнения функции
func logParseUpdateBoolUp(id int, col, val string, colVals map[string]string,
	ctx context.Context, db model.DBLog, w http.ResponseWriter) error {

	sst, err := strconv.ParseBool(colVals[strings.ToUpper(col)])
	if err != nil {
		return err
	}

	err = UpdateContext(ctx, db, col, sst, id)
	if err != nil {
		return err
	}

	return nil
}

// преобразование строки к Bool, проверка ошибки, выполнение patch-запроса
// и сообщение об успешности выполнения функции
func logParseUpdateBoolLow(id int, col, val string,
	ctx context.Context, db model.DBLog, w http.ResponseWriter) error {

	sst, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}

	err = UpdateContext(ctx, db, col, sst, id)
	if err != nil {
		return err
	}

	return nil
}

func scanRow(r *http.Request) map[string]string {
	return map[string]string{
		"id":            r.URL.Query().Get("id"),
		"auth":          r.URL.Query().Get("auth"),
		"ip":            r.URL.Query().Get("ip"),
		"stream":        r.URL.Query().Get("stream"),
		"run":           r.URL.Query().Get("run"),
		"portsrv":       r.URL.Query().Get("portsrv"),
		"sp":            r.URL.Query().Get("sp"),
		"camid":         r.URL.Query().Get("camid"),
		"record_status": r.URL.Query().Get("record_status"),
		"stream_status": r.URL.Query().Get("stream_status"),
		"record_state":  r.URL.Query().Get("record_state"),
		"stream_state":  r.URL.Query().Get("stream_state"),
	}
}
