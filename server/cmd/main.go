package main

import (
	"context"
	"server/config"
	"server/db"
	"server/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			config.LoadConf,
			db.NewDatabase,
			server.NewHttpServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)

	app.Run()
}
