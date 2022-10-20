package main

import (
	"context"

	"github.com/Kseniya-cha/server/methods"
	"github.com/Kseniya-cha/server/model"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg model.Config

func main() {
	// чтение конфига
	errRC := cleanenv.ReadConfig("config.yml", &cfg)

	// подключение к базе, инициализация логас уровнем логирования LogLevel, указанным в
	// конфиге (структура DbLog) close - функция для отложенного закрытия файла .log
	DbLog, close := model.NewDBLog(cfg, cfg.LogLevel)
	// отложенное закрытие файла .log
	defer close()
	// отложенное закрытие базы данных
	defer cfg.CloseDB(DbLog)
	// обработка ошибки чтения конфига
	DbLog.LogIFAction(errRC, "read config.yml")

	// инициализация лога с уровнем логирования, указанным в конфиге
	// log, close := model.DefLog(cfg.LogLevel)

	ctx := context.Background()

	// корректное завершение работы программы
	// при генерации сигнала
	go methods.GracefulShutdown(DbLog.Log)

	// инициализация router
	router := mux.NewRouter()

	// добавление урлов с логированием
	// http://localhost:3333/
	hfRoot := methods.RootHF(DbLog)
	router.HandleFunc("/", hfRoot)

	// http://localhost:3333/api/get/
	hfSelect := methods.GetAllHF(ctx, DbLog, true)
	router.HandleFunc("/api/get/", hfSelect).Methods("GET")

	// http://localhost:3333/api/get/3/
	hfGetId := methods.GetIdHF(ctx, DbLog, true)
	router.HandleFunc("/api/get/{ID}/", hfGetId).Methods("GET")

	// 'http://localhost:3333/api/delete/3/'
	hfDeleteId := methods.DeleteIdHF(ctx, DbLog)
	router.HandleFunc("/api/delete/{ID}/", hfDeleteId).Methods("DELETE")

	// http://localhost:3333/api/post/auth/ip/stream/run/port/sp/cam/true/false/false/true/
	hfPost := methods.PostHF(ctx, DbLog)
	router.HandleFunc("/api/post/{AUTH}/{IP}/{STREAM}/{RUN}/{PORTSRV}/{SP}/{CAMID}/{RECORD_STATUS}/{STREAM_STATUS}/{RECORD_STATE}/{STREAM_STATE}/",
		hfPost).Methods("POST")

	// http://localhost:3333/api/put/6/auth/ip/stream/run/port/sp/cam/true/false/false/true/
	hfPut := methods.PutHF(ctx, DbLog)
	router.HandleFunc("/api/put/{ID}/{AUTH}/{IP}/{STREAM}/{RUN}/{PORTSRV}/{SP}/{CAMID}/{RECORD_STATUS}/{STREAM_STATUS}/{RECORD_STATE}/{STREAM_STATE}/",
		hfPut).Methods("PUT")

	// http://localhost:3333/api/patch
	//?id=&auth=&
	hfPatch := methods.PatchHF(ctx, DbLog)
	router.HandleFunc("/api/patch", hfPatch).Methods("PATCH")

	// инициализации сервера
	server := cfg.NewServer(router)

	// запуск сервера
	methods.RunServer(DbLog, server)
}
