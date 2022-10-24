package main

import (
	refreshstream "github.com/Kseniya-cha/server/internal/refreshStream"
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

	// корректное завершение работы программы при генерации сигнала
	go dblog.GracefulShutdown(DbLog)

	// проверка подключения к базе данных
	go dblog.PingDB(DbLog, cfg)

	// инициализация router
	router := mux.NewRouter()

	// добавление урлов с логированием
	refreshstream.RegisterRouter(router, DbLog)

	// инициализации сервера
	serv := server.NewServer(cfg, router)

	// запуск сервера
	server.RunServer(DbLog, serv)
}
