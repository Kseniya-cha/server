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

	fmt.Println(cfg)

	// инициализация сервера
	app := server.NewApp(cfg)

	go app.GracefulShutdown()

	if err := app.RunServer(); err != nil {
		fmt.Println(err)
	}
}
