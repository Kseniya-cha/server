package main

import (
	"context"

	"github.com/Kseniya-cha/server/logger"
	"github.com/Kseniya-cha/server/methods"
	"github.com/Kseniya-cha/server/model"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg model.Config

func main() {
	// чтение конфига
	errRC := cleanenv.ReadConfig("config.yml", &cfg)

	// инициализация лога с уровнем логирования, указанным в конфиге
	log, close := logger.DefLog(cfg.LogLevel)
	// отложенное закрытие файла .log
	defer close()
	// обработка ошибки чтения конфига
	logger.LogIFAction(log, errRC, "read config.yml")

	ctx := context.Background()

	// подключение к базе данных и отложенное закрытие
	db := cfg.OpenDB(log)
	defer cfg.CloseDB(db, log)

	// корректное завершение работы программы
	// при генерации сигнала
	go methods.GracefulShutdown(log)

	// инициализация router
	router := mux.NewRouter()

	// добавление урлов с логированием
	hfRoot := methods.RootHF(log)
	router.HandleFunc("/", hfRoot)

	// curl 'http://localhost:3333/api/smth/3/aaa/bbb'
	hfSmth := methods.SmthHF()
	router.HandleFunc("/api/smth/{ID}/{IP}/{SMTH}/", hfSmth).Methods("GET")

	// curl http://localhost:3333/api/get/
	hfSelect := methods.GetAllHF(ctx, db, log, true)
	router.HandleFunc("/api/get/", hfSelect).Methods("GET")
	// curl http://localhost:3333/api/get/3/
	hfGetId := methods.GetIdHF(ctx, db, log, true)
	router.HandleFunc("/api/get/{ID}/", hfGetId).Methods("GET")
	// curl 'http://localhost:3333/api/delete/3/'
	hfDeleteId := methods.DeleteIdHF(ctx, db, log)
	router.HandleFunc("/api/delete/{ID}/", hfDeleteId).Methods("DELETE")
	// curl http://localhost:3333/api/put/6/auth/ip/stream/run/port/sp/cam/true/false/false/true/
	hfPut := methods.PutHF(ctx, db, log)
	router.HandleFunc("/api/put/{ID}/{AUTH}/{IP}/{STREAM}/{RUN}/{PORTSRV}/{SP}/{CAMID}/{RECORD_STATUS}/{STREAM_STATUS}/{RECORD_STATE}/{STREAM_STATE}/",
		hfPut).Methods("PUT")
	hfPost := methods.PostHF(ctx, db, log)
	router.HandleFunc("/api/post/{ID}/", hfPost).Methods("POST")

	// инициализации сервера
	server := cfg.NewServer(router)

	// запуск сервера
	methods.RunServer(log, server)
}
