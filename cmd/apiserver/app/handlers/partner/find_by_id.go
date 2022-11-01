package partner

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PartnerHandler) FindByID(ctx *gin.Context) {
	id := uint(1)
	h.PartnerService.FindByID(ctx, &id)

	ctx.JSON(http.StatusOK, api.Message{
		Message: "",
	})
}
