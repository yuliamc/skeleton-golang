package partner

import (
	"modalrakyat/skeleton-golang/internal/services/partner"
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/messages"
	"modalrakyat/skeleton-golang/pkg/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *PartnerHandler) PostCreate(ctx *gin.Context) {
	payload := &partner.CreatePartnerPayload{}
	if err := ctx.ShouldBindWith(payload, binding.JSON); err != nil {
		errors.ResponseErrorWithString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(payload); err != nil {
		errors.ResponseErrorWithString(ctx, errorMessage)
		return
	}
	if err := h.PartnerService.Create(ctx, payload); err != nil {
		errors.ResponseError(ctx, err)
	} else {
		ctx.JSON(http.StatusCreated, api.Message{
			Message: messages.TranslateCode(ctx, int(messages.MSG_CODE_INSERT_SUCCESS)),
		})
	}
}
