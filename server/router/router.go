package router

import (
	"net/http"
	"server/config"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, config *config.Configuration) {

	routes := r.Group("/api/v1")

	routes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "chat-server", "version": "1.0.0"})
	})
}
