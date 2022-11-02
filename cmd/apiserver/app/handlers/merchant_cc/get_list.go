package merchant_cc

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *MerchantCCHandler) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.BaseWithMeta{
		Data: []struct {
			CCNumber  string  `json:"va_number"`
			Provider  string  `json:"provider"`
			RefID     string  `json:"ref_id"`
			Limit     float64 `json:"limit"`
			ExpiredAt string  `json:"expired_at"`
		}{
			{CCNumber: "371449635398431", Provider: "Visa", RefID: "76009244561", Limit: 50000000, ExpiredAt: "2023-03"},
			{CCNumber: "3530111333300000", Provider: "MasterCard", RefID: "5019717010103742", Limit: 30000000, ExpiredAt: "2025-11"},
			{CCNumber: "5105105105105100", Provider: "JCB", RefID: "6331101999990016", Limit: 15000000, ExpiredAt: "2024-06"},
			{CCNumber: "4012888888881881", Provider: "American Express", RefID: "4222222222222", Limit: 8000000, ExpiredAt: "2022-12"},
		},
		Meta: map[string]interface{}{
			"total": 2,
		},
	})
}
