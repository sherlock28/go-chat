package main

import (
	"context"
	"server/config"
	"server/db"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			config.LoadConf,
			db.NewDatabase,
		),
		fx.Invoke(func(*db.Database) {}),
	)

	app.Run()
}
