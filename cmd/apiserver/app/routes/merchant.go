package routes

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/store"
	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func initMerchantRoute(group *gin.RouterGroup) {
	group.Use(middlewares.StaticApiKey(&config.Config.SecretKey.StaticApiKey))

	// CC.
	group.GET("/va", store.MerchantCCHandler.GetList)

	// VA.
	group.GET("/cc", store.MerchantVAHandler.GetList)
}
