package partner

import (
	"modalrakyat/skeleton-golang/internal/services/partner"
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/null"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PartnerHandler) GetByID(ctx *gin.Context) {
	uri := &partner.Uri{}
	if err := ctx.ShouldBindUri(uri); err != nil {
		errors.ResponseError(ctx, err)
		return
	}

	if partner, err := h.PartnerService.FindByID(ctx, &uri.PartnerID); err != nil {
		errors.ResponseError(ctx, err)
	} else {
		if null.IsNil(partner) {
			errors.ResponseError(ctx, errors.NewGenericError(int(errors.ERROR_MSG_DATA_NOT_FOUND)))
			return
		}

		ctx.JSON(http.StatusOK, api.Base{
			Data: partner,
		})
	}
}
