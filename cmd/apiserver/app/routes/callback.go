package routes

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/store"

	"github.com/gin-gonic/gin"
)

func initCallbackRoute(group *gin.RouterGroup) {
	group.POST("/bca/notification", store.CallbackHandler.PostBCANotification)
	group.POST("/visa/notification", store.CallbackHandler.PostVisaNotification)
}
