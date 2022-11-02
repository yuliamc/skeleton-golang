package callback

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CallbackHandler) PostVisaNotification(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, api.Message{
		Message: messages.TranslateCode(ctx, int(messages.MSG_CODE_NOTIFICATION_ACCEPTED)),
	})
}
