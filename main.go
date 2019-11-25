package main

import (
	log "todo-graph/logger"
	"todo-graph/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			server.New,
			log.New,
			server.GetAllHandlers,
		),
		fx.Invoke(server.RegisterHanders, server.InitServer),
	)
	app.Run()
}
