package bo_config

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func (h *BOConfigHandler) GetConfig(ctx *gin.Context) {

	const CACHE_KEY = "bo_admin:config"

	config := make(map[string]interface{}, 0)
	if err := h.redisClient.Get(CACHE_KEY, &config); err != nil {
		if err == redis.Nil {
			// do nothing, it means key not found on redis storage
		} else {
			errors.ResponseError(ctx, errors.NewGenericError(int(errors.ERROR_MSG_INTERNAL_SERVER_ERROR)))
			return
		}
	}
	ctx.JSON(http.StatusOK, api.Base{
		Data: config,
	})
}
