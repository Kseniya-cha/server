package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type refreshStreamHandler struct {
	db      *sql.DB
	log     *logrus.Logger
	useCase refreshStream.RefreshStreamUseCase
}

func NewRefreshStreamHandler(useCase refreshStream.RefreshStreamUseCase, db *sql.DB, log *logrus.Logger) *refreshStreamHandler {
	return &refreshStreamHandler{
		db:      db,
		log:     log,
		useCase: useCase,
	}
}

// http://localhost:3333/api/get/
func (s *refreshStreamHandler) GetAllHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// запрос
		data, err := s.useCase.Get(ctx)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// вывод данных
		logger.LogWriteDebug(s.log, w, fmt.Sprintf("%v", data))

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, refreshStream.GetHFRespOkConst)
	}
}

// показать все строки по значению id
// http://localhost:3333/api/get/3/
func (s *refreshStreamHandler) GetIdHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// извлечение и парсинг id
		value := mux.Vars(r)["ID"]
		val, err := strconv.Atoi(value)
		if err != nil {
			logger.LogError(s.log, err)
		} else {
			logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
		}

		// запрос
		data, err := s.useCase.GetId(ctx, val)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// вывод данных
		logger.LogWriteDebug(s.log, w, fmt.Sprintf("%v", data))

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.GetIdHFRespOkConst, val))
	}
}

// http://localhost:3333/api/delete/3/
func (s *refreshStreamHandler) DeleteIdHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// извлечение и парсинг id
		value := mux.Vars(r)["ID"]
		val, err := strconv.Atoi(value)
		if err != nil {
			logger.LogError(s.log, err)
		} else {
			logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
		}

		// запрос
		err = s.useCase.Delete(ctx, val)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.DeleteHFRespOkConst, val))
	}
}

// вставить строку с помощью json-файла
func (s *refreshStreamHandler) PostHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)

		allcols := refreshStream.PostHFAllColsConst
		allvalues := fmt.Sprintf(refreshStream.PostHFAllValuesConst,
			rs.Auth.String, rs.Ip.String, rs.Stream.String, rs.Run.String,
			rs.Portsrv, rs.Sp.String, rs.Camid.String, rs.Record_status.Bool,
			rs.Stream_status.Bool, rs.Record_state.Bool, rs.Stream_state.Bool)

		// вставка данных и обработка ошибки
		err = s.useCase.Insert(ctx, allcols, allvalues)
		if err != nil {
			logger.LogDebug(s.log, err)
		}

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PostHFRespOkConst))
	}
}

// изменение строки
func (s *refreshStreamHandler) PutHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)

		// извлечение id
		id := rs.Id
		if id == 0 {
			logger.LogError(s.log, "this Id does not exist!")
			return
		}

		// выолнение запроса
		err = s.useCase.Update(ctx, "auth", rs.Auth.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "ip", rs.Ip.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream", rs.Stream.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "run", rs.Run.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "portsrv", rs.Portsrv, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "sp", rs.Sp.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "camid", rs.Camid.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "record_status", rs.Record_status.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream_status", rs.Stream_status.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "record_state", rs.Record_state.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream_state", rs.Stream_state.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PostHFRespOkConst))
	}
}

// частичное изменение строки по id
func (s *refreshStreamHandler) PatchHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)

		id := rs.Id
		if id == 0 {
			logger.LogError(s.log, "this Id does not exist!")
			return
		}

		err = s.useCase.Update(ctx, "auth", rs.Auth.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "ip", rs.Ip.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream", rs.Stream.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "run", rs.Run.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "portsrv", rs.Portsrv, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "sp", rs.Sp.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "camid", rs.Camid.String, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "record_status", rs.Record_status.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream_status", rs.Stream_status.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "record_state", rs.Record_state.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		err = s.useCase.Update(ctx, "stream_state", rs.Stream_state.Bool, id)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// сообщение о завершении (добавить код ошибки!!)
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PostHFRespOkConst))
	}
}
