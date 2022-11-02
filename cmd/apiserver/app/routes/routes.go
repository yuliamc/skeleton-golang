package routes

import (
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

	// Setup pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	initMerchantRoute(r.Group("/merchant"))
	initBackofficeRoute(r.Group("/bo"))
	initCallbackRoute(r.Group("/callback"))

	return r
}
