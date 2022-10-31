package partner

type PartnerResponse struct {
	ID                 uint                     `json:"id" binding:"required"`
	Code               string                   `json:"code" binding:"required"`
	UniqueID           string                   `json:"unique_id" binding:"required"`
	Name               string                   `json:"name" binding:"required"`
	LegalName          string                   `json:"legal_name" binding:"required"`
	LoanTypeID         uint                     `json:"loan_type_id" binding:"required"`
	LoanTypeName       string                   `json:"loan_type_name" binding:"required"`
	LoanCodePrefix     string                   `json:"loan_code_prefix" binding:"required"`
	InitialCreditLimit float64                  `json:"initial_credit_limit" binding:"required"`
	Settings           []map[string]interface{} `json:"settings" binding:"required"`
	LoanScheme         []map[string]interface{} `json:"loan_scheme" binding:"required"`
	CustomSetting      map[string]interface{}   `json:"custom_setting" binding:"required"`
	EscrowAccount      map[string]interface{}   `json:"escrow_account" binding:"required"`
}

type PartnerDropdownResponse struct {
	Data *[]PartnerDropdownItemResponse `json:"data"`
	Meta struct {
		TotalData int `json:"total_data"`
	} `json:"meta"`
}

type PartnerDropdownItemResponse struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
