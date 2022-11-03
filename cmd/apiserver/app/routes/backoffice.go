package routes

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/store"

	"github.com/gin-gonic/gin"
)

func initBackofficeRoute(group *gin.RouterGroup) {
	// Auth.
	group.POST("/auth/signin", store.BOAuthHandler.PostSignin)

	// After this line, all endpoints require login.
	group.Use(store.BackofficeAuthMiddleware.Authenticate())

	// Admin.
	group.GET("/admin", store.BOAdminHandler.GetList)

	// Partner.
	group.GET("/partner", store.PartnerHandler.GetList)
	group.GET("/partner-detail/:id", store.PartnerHandler.GetByID)
	group.POST("/partner", store.PartnerHandler.PostCreate)

	// Config.
	group.GET("/config", store.BOConfigHandler.GetConfig)
	group.POST("/config", store.BOConfigHandler.PostConfig)
	group.DELETE("/config", store.BOConfigHandler.DeleteConfig)
}
