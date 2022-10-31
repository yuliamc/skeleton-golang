package loan_scheme

import (
	"modalrakyat/skeleton-golang/internal/services/partner"
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (h *LoanSchemeHandler) GetLoanScheme(ctx *gin.Context) {
	// Gets partner from middleware.
	partnerCtx, exist := ctx.Get("partner")
	if !exist {
		errors.ErrorCode(ctx, http.StatusBadRequest, errors.CLIENT_INFO_NOT_FOUND)
		return
	}

	partner := partnerCtx.(*partner.PartnerResponse)

	var minNominalLoan []float64
	var maxNominalLoan []float64
	var minTenor []float64
	var maxTenor []float64
	for _, v := range partner.Settings {
		minNominalLoan = append(minNominalLoan, v["min_nominal_loan"].(float64))
		maxNominalLoan = append(maxNominalLoan, v["max_nominal_loan"].(float64))
		minTenor = append(minTenor, v["min_tenor"].(float64))
		maxTenor = append(maxTenor, v["max_tenor"].(float64))
	}
	sort.Float64s(minNominalLoan)
	sort.Float64s(maxNominalLoan)
	sort.Float64s(minTenor)
	sort.Float64s(maxTenor)

	ctx.JSON(http.StatusCreated, api.Base{
		Data: response{
			MinNominalLoan: minNominalLoan[0],
			MaxNominalLoan: maxNominalLoan[len(maxNominalLoan)-1],
			MinTenor:       minTenor[0],
			MaxTenor:       maxTenor[len(maxTenor)-1],
		},
	})
}

type response struct {
	MinNominalLoan float64 `json:"min_nominal_loan"`
	MaxNominalLoan float64 `json:"max_nominal_loan"`
	MinTenor       float64 `json:"min_tenor"`
	MaxTenor       float64 `json:"max_tenor"`
}
