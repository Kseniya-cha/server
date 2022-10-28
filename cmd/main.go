package main

import (
	"fmt"

	"github.com/Kseniya-cha/server/pkg/config"
	"github.com/Kseniya-cha/server/pkg/server"
)

// var cfg *config.Config

func main() {
	// чтение конфига
	cfg := config.GetConfig()

	// инициализация сервера
	app := server.NewApp(cfg)

	// ожидание прерывающего сигнала
	// в отдельной горутине
	go app.GracefulShutdown()

	// запуск сервера
	if err := app.RunServer(); err != nil {
		fmt.Println(err)
	}
}
