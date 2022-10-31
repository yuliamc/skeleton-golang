package loan_scheme

import (
	srvs "modalrakyat/skeleton-golang/internal/services"
	"modalrakyat/skeleton-golang/internal/services/partner"
)

type LoanSchemeHandler struct {
	RegistrationPartnerService srvs.RegistrationPartnerService
	LoanSchemeService          srvs.LoanSchemeService
}

func NewLoanSchemeHandler(RegistrationPartnerService srvs.RegistrationPartnerService, LoanSchemeService srvs.LoanSchemeService) *LoanSchemeHandler {
	return &LoanSchemeHandler{
		RegistrationPartnerService: RegistrationPartnerService,
		LoanSchemeService:          LoanSchemeService,
	}
}

// Find match setting by nominal loan and tenor from partner settings.
func findMatchedSetting(nominalLoan *float64, tenor *int, partner *partner.PartnerResponse) map[string]interface{} {
	if partner.Settings == nil || len(partner.Settings) == 0 {
		return nil
	}

	var foundSetting map[string]interface{}
	for _, v := range partner.Settings {
		settingMinNominalLoan, ok := v["min_nominal_loan"].(float64)
		if !ok {
			continue
		}
		settingMaxNominalLoan, ok := v["max_nominal_loan"].(float64)
		if !ok {
			continue
		}
		settingMinTenor, ok := v["min_tenor"].(float64)
		if !ok {
			continue
		}
		settingMaxTenor, ok := v["max_tenor"].(float64)
		if !ok {
			continue
		}
		if !(*nominalLoan >= settingMinNominalLoan &&
			*nominalLoan <= settingMaxNominalLoan &&
			*tenor >= int(settingMinTenor) &&
			*tenor <= int(settingMaxTenor)) {
			continue
		}

		foundSetting = v
		break
	}

	return foundSetting
}
