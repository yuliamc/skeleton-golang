package bo_config

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/messages"
	"modalrakyat/skeleton-golang/pkg/utils/net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * Example configuration JSON body
 * { "is_enabled": true, "va_bank_id": 1234, "allowed_cc_bank_id": [ 1234 ] }
 */
func (h *BOConfigHandler) PostConfig(ctx *gin.Context) {

	const CACHE_KEY = "bo_admin:config"

	requestParser := net.HTTPRequestParser{}
	requestParser.ParseRequestBody(ctx)
	data := requestParser.GetRequestBody(ctx)

	if err := h.redisClient.Set(CACHE_KEY, data, 5*time.Minute); err != nil {
		errors.ResponseError(ctx, err)
		return
	} else {
		ctx.JSON(http.StatusCreated, api.Message{
			Message: messages.TranslateCode(ctx, int(messages.MSG_CODE_INSERT_SUCCESS)),
		})
	}
}
