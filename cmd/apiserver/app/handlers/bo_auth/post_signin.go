package bo_auth

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BOAuthHandler) PostSignin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.Base{
		Data: struct {
			SID       string `json:"sid"`
			ExpiredAt string `json:"expired_at"`
		}{
			SID:       "12345",
			ExpiredAt: "2023-31-12 23:59:59",
		},
	})
}
