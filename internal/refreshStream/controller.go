package refreshstream

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Kseniya-cha/server/constants"
	dblog "github.com/Kseniya-cha/server/pkg/DBLog"
	"github.com/gorilla/mux"
)

// http://localhost:3333/
func RootHF(db dblog.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db.LogI("get page \"/\"")
		fmt.Printf(constants.RootPageConst)
		io.WriteString(w, constants.RootPageConst)
	}
}

// http://localhost:3333/api/get/
func GetAllHF(ctx context.Context, db dblog.DBLog,
	show bool) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		jData, err := SelectContextRS(ctx, db)
		db.LogPrintFat(err)

		if show {
			db.LogWrite(w, fmt.Sprintf(string(jData)))
		}
		db.LogWrite(w, constants.GetAllHFRespOkConst)
	}
}

// показать все строки по значению id
// http://localhost:3333/api/get/3/
func GetIdHF(ctx context.Context, db dblog.DBLog,
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
		db.LogWrite(w, fmt.Sprintf(constants.GetIdHFRespOkConst, val))
	}
}

// http://localhost:3333/api/delete/3/
func DeleteIdHF(ctx context.Context, db dblog.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		value := mux.Vars(r)["ID"]
		val, err := strconv.Atoi(value)
		db.LogIFAction(err, constants.ConvIdIntConst)

		err = DeleteContext(ctx, db, val)
		db.LogWriteIF(w, err, "delete")
	}
}

// вставить строку с помощью json-файла
func PostHFJSON(ctx context.Context, db dblog.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs RefreshStreamWithNull
		err := decoder.Decode(&rs)
		db.LogWriteIF(w, err, "decode json")

		allcols := constants.PostHFAllColsConst
		allvalues := fmt.Sprintf(constants.PostHFAllValuesConst,
			rs.Auth.String, rs.Ip.String, rs.Stream.String, rs.Run.String,
			rs.Portsrv, rs.Sp.String, rs.Camid.String, rs.Record_status.Bool,
			rs.Stream_status.Bool, rs.Record_state.Bool, rs.Stream_state.Bool)

		// вставка данных и обработка ошибки
		err = InsertContext(ctx, db, allcols, allvalues)
		db.LogWriteIF(w, err, "insert")
	}
}

// вставить строку
func PutHFJSON(ctx context.Context, db dblog.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs RefreshStreamWithNull
		err := decoder.Decode(&rs)
		db.LogWriteIF(w, err, "decode json")

		id := rs.Id

		err = UpdateContext(ctx, db, "auth", rs.Auth.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "ip", rs.Ip.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream", rs.Stream.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "run", rs.Run.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "portsrv", rs.Portsrv, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "sp", rs.Sp.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "camid", rs.Camid.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "record_status", rs.Record_status.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream_status", rs.Stream_status.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "record_state", rs.Record_state.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream_state", rs.Stream_state.Bool, id)
		db.LogPrintFat(err)

	}
}

// частичное изменение строки по id
func PatchHFJSON(ctx context.Context, db dblog.DBLog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs RefreshStreamWithNull
		err := decoder.Decode(&rs)
		db.LogWriteIF(w, err, "decode json")

		id := rs.Id

		err = UpdateContext(ctx, db, "auth", rs.Auth.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "ip", rs.Ip.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream", rs.Stream.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "run", rs.Run.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "portsrv", rs.Portsrv, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "sp", rs.Sp.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "camid", rs.Camid.String, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "record_status", rs.Record_status.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream_status", rs.Stream_status.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "record_state", rs.Record_state.Bool, id)
		db.LogPrintFat(err)
		err = UpdateContext(ctx, db, "stream_state", rs.Stream_state.Bool, id)
		db.LogPrintFat(err)
	}
}

func RegisterRouter(router *mux.Router, db dblog.DBLog) {
	ctx := context.Background()

	// http://localhost:3333/
	hfRoot := RootHF(db)
	router.HandleFunc(constants.URLRootConst, hfRoot)

	// http://localhost:3333/api/get/
	hfSelect := GetAllHF(ctx, db, true)
	router.HandleFunc(constants.URLApiConst, hfSelect).Methods("GET")

	// http://localhost:3333/api/get/3/
	hfGetId := GetIdHF(ctx, db, true)
	router.HandleFunc(constants.URLGetDelIdConst, hfGetId).Methods("GET")

	// http://localhost:3333/api/delete/3/
	hfDeleteId := DeleteIdHF(ctx, db)
	router.HandleFunc(constants.URLGetDelIdConst, hfDeleteId).Methods("DELETE")

	hfPostJS := PostHFJSON(ctx, db)
	router.HandleFunc(constants.URLApiConst,
		hfPostJS).Methods("POST")

	hfPutJS := PutHFJSON(ctx, db)
	router.HandleFunc(constants.URLApiConst,
		hfPutJS).Methods("PUT")

	hfPatchJS := PatchHFJSON(ctx, db)
	router.HandleFunc(constants.URLApiConst,
		hfPatchJS).Methods("PATCH")
}
