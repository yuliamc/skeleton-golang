package partner

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PartnerHandler) GetList(ctx *gin.Context) {
	if partnerResponses, count, err := h.PartnerService.FindAll(ctx); err != nil {
		errors.ResponseError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, api.BaseWithMeta{
			Data: partnerResponses,
			Meta: map[string]interface{}{
				"total": count,
			},
		})
	}
}
