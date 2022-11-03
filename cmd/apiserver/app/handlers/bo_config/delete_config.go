package bo_config

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BOConfigHandler) DeleteConfig(ctx *gin.Context) {

	const CACHE_KEY = "bo_admin:config"

	if err := h.redisClient.Delete(CACHE_KEY); err != nil {
		errors.ResponseError(ctx, errors.NewGenericError(int(errors.ERROR_MSG_INTERNAL_SERVER_ERROR)))
		return
	}
	ctx.JSON(http.StatusOK, api.Message{
		Message: messages.TranslateCode(ctx, int(messages.MSG_CODE_DELETE_SUCCESS)),
	})
}
