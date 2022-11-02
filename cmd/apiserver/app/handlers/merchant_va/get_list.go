package merchant_va

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *MerchantVAHandler) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.BaseWithMeta{
		Data: []struct {
			VANumber string  `json:"va_number"`
			Provider string  `json:"provider"`
			RefID    string  `json:"ref_id"`
			Amount   float64 `json:"amount"`
			Type     string  `json:"type"`
			Status   string  `json:"status"`
		}{
			{VANumber: "12054991000004070", Provider: "BCA", RefID: "1e6a7ef0-18af-11e9-9b8c-8bb266ad7db6", Amount: 0, Type: "open_amount", Status: "open"},
			{VANumber: "9881001910007393", Provider: "BNI", RefID: "RPY018112204395502", Amount: 187572, Type: "close_amount", Status: "closed"},
			{VANumber: "9881001910007204", Provider: "BNI", RefID: "RPY2102040939350001", Amount: 9114442, Type: "close_amount", Status: "expired"},
			{VANumber: "1293602216009342", Provider: "BCA", RefID: "eacf3cf0-10d5-11ea-9d4f-731bd8ab65ad", Amount: 0, Type: "open_amount", Status: "open"},
		},
		Meta: map[string]interface{}{
			"total": 2,
		},
	})
}
