package main

import (
	"context"

	"github.com/Kseniya-cha/server/constants"
	"github.com/Kseniya-cha/server/internal/refreshStream"
	dblog "github.com/Kseniya-cha/server/pkg/DBLog"
	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/Kseniya-cha/server/pkg/server"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func main() {
	// чтение конфига
	errRC := cleanenv.ReadConfig("config.yml", &cfg)

	// подключение к базе, инициализация лога(структура DbLog)
	// close - функция для отложенного закрытия файла out.log
	DbLog, close := dblog.NewDBLog(cfg)

	// отложенное закрытие файла .log и базы данных
	defer close()
	defer dblog.CloseDB(DbLog, cfg)

	// обработка ошибки чтения конфига
	DbLog.LogIFAction(errRC, "read config.yml")

	ctx := context.Background()

	// корректное завершение работы программы при генерации сигнала
	go dblog.GracefulShutdown(DbLog)

	// проверка подключения к базе данных
	go dblog.PingDB(DbLog, cfg)

	// инициализация router
	router := mux.NewRouter()

	// добавление урлов с логированием
	// http://localhost:3333/
	hfRoot := refreshstream.RootHF(DbLog)
	router.HandleFunc(constants.URLRootConst, hfRoot)

	// http://localhost:3333/api/get/
	hfSelect := refreshstream.GetAllHF(ctx, DbLog, true)
	router.HandleFunc(constants.URLGetAllConst, hfSelect).Methods("GET")

	// http://localhost:3333/api/get/3/
	hfGetId := refreshstream.GetIdHF(ctx, DbLog, true)
	router.HandleFunc(constants.URLGetIdConst, hfGetId).Methods("GET")

	// 'http://localhost:3333/api/delete/3/'
	hfDeleteId := refreshstream.DeleteIdHF(ctx, DbLog)
	router.HandleFunc(constants.URLDeleteIdConst, hfDeleteId).Methods("DELETE")

	// http://localhost:3333/api/post/auth/ip/stream/run/port/sp/cam/true/false/false/true/
	hfPost := refreshstream.PostHF(ctx, DbLog)
	router.HandleFunc(constants.URLPostConst,
		hfPost).Methods("POST")

	// http://localhost:3333/api/put/6/auth/ip/stream/run/port/sp/cam/true/false/false/true/
	hfPut := refreshstream.PutHF(ctx, DbLog)
	router.HandleFunc(constants.URLPutConst,
		hfPut).Methods("PUT")

	// http://localhost:3333/api/patch
	//?id=&auth=&
	hfPatch := refreshstream.PatchHF(ctx, DbLog)
	router.HandleFunc(constants.URLPatchConst, hfPatch).Methods("PATCH")

	// инициализации сервера
	serv := server.NewServer(cfg, router)

	// запуск сервера
	server.RunServer(DbLog, serv)
}
