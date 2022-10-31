package routes

import (
	"modalrakyat/skeleton-golang/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// Init initialize application routes
// DO NOT FORGET TO READ README#routes FIRST!
func Init(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// Setup middlewares
	r.Use(middlewares.Recovery(mode))
	// r.Use(middlewares.RequestID())
	r.Use(middlewares.AccessLog())
	r.Use(middlewares.LanguageAccept())
	r.Use(middlewares.Cors())
	// r.Use(middlewares.LogPath)

	// Setup pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Setup application route
	appsRouteGroup := r.Group("/apps")
	// appsRouteGroup.Use(store.MiddlewareAccess.ClientAuthentication())
	// appsRouteGroup.Use(store.MiddlewareAccess.WhitelistIP())
	initApplicationRoute(appsRouteGroup)

	return r
}
