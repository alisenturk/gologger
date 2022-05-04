package main

import (
	"github.com/alisenturk/gologger/config"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfig()

	server := router.NewServer(api, config, log)
	app := App{config, server}
	log.Info("App Server Starting...")
	app.run(log)

}

type App struct {
	Config *config.Config
	Server *router.Echo
}

func (a *App) run(log *zap.Logger) {
	a.Server.RunEcho()
}
