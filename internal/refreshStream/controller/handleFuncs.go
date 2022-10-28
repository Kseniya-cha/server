package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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
func (s *refreshStreamHandler) GetHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// запрос
		data, err := s.useCase.Get(ctx)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// вывод данных
		logger.LogWriteDebug(s.log, w, fmt.Sprintf("%v", data))

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.GetHFRespOkConst, http.StatusOK))
	}
}

// показать все строки по значению id
// http://localhost:3333/api/get/3/
func (s *refreshStreamHandler) GetIdHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// извлечение
		value := mux.Vars(r)["ID"]
		// val, err := strconv.Atoi(value)
		// if err != nil {
		// 	logger.LogError(s.log, err)
		// } else {
		// 	logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
		// }

		// запрос
		data, err := s.useCase.GetId(ctx, value)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// вывод данных
		logger.LogWriteDebug(s.log, w, fmt.Sprintf("%v", data))

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.GetIdHFRespOkConst, value, http.StatusOK))
	}
}

// http://localhost:3333/api/delete/3/
func (s *refreshStreamHandler) DeleteIdHF(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// извлечение и парсинг id
		value := mux.Vars(r)["ID"]
		// val, err := strconv.Atoi(value)
		// if err != nil {
		// 	logger.LogError(s.log, err)
		// } else {
		// 	logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
		// }

		// запрос
		err := s.useCase.Delete(ctx, value)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.DeleteHFRespOkConst, value, http.StatusOK))
	}
}

// вставить строку с помощью json-файла
func (s *refreshStreamHandler) PostHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs *refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)

		// вставка данных и обработка ошибки
		err = s.useCase.Insert(ctx, rs)
		if err != nil {
			logger.LogDebug(s.log, err)
		}

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PostHFRespOkConst, http.StatusOK))
	}
}

// изменение строки
func (s *refreshStreamHandler) PutHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs *refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)

		// выполнение запроса
		err = s.useCase.Update(ctx, rs)
		if err != nil {
			logger.LogError(s.log, fmt.Sprintf(refreshStream.PutRespErrColConst, err))
			return
		}

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PutHFRespOkConst, http.StatusOK))
	}
}

// частичное изменение строки по id
func (s *refreshStreamHandler) PatchHFJSON(ctx context.Context) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// преобразование полученного json-файла и обработка ошибки
		decoder := json.NewDecoder(r.Body)
		var rs *refreshStream.RefreshStreamWithNull
		err := decoder.Decode(&rs)
		if err != nil {
			logger.LogError(s.log, err)
			return
		}
		logger.LogDebug(s.log, refreshStream.DecodeJsonConst)
		fmt.Println(rs)

		// выполнение запроса
		err = s.useCase.Update(ctx, rs)
		if err != nil {
			logger.LogError(s.log, fmt.Sprintf(refreshStream.PutRespErrColConst, err))
			return
		}

		// сообщение о завершении
		logger.LogWriteInfo(s.log, w, fmt.Sprintf(refreshStream.PatchHFRespOkConst, http.StatusOK))
	}
}
