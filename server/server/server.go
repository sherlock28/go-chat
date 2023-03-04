package server

import (
	"context"
	"log"
	"server/config"
	"server/router"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, config *config.Configuration) *gin.Engine {
	log.Println("⌛ Starting server...")

	r := gin.Default()

	r.Use(cors.New(GetCORSConfig()))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.InitRouter(r, config)
			go r.Run("0.0.0.0:" + strconv.Itoa(config.App.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return r
}

func GetCORSConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}
}
