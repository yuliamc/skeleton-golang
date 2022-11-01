package routes

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/store"

	"github.com/gin-gonic/gin"
)

func initApplicationRoute(group *gin.RouterGroup) {
	group.GET("/application/partner-detail/:id", store.PartnerHandler.GetByID)
}
