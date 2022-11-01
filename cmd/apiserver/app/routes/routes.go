package routes

import (
	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// Setup middlewares
	r.Use(middlewares.Recovery(mode))
	r.Use(middlewares.AccessLog())
	r.Use(middlewares.LanguageAccept())
	r.Use(middlewares.Cors())
	r.Use(middlewares.StaticApiKey(&config.Config.SecretKey.StaticApiKey))

	// Setup pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Setup application route
	appsRouteGroup := r.Group("/apps")
	initApplicationRoute(appsRouteGroup)

	return r
}
