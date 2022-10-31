package partner

import (
	"modalrakyat/skeleton-golang/internal/model"
)

func (c *PartnerResponse) PartnerStructResponse(data model.Partner) {
	c.ID = data.ID
	c.UniqueID = data.UniqueID
	c.Code = data.Code
	c.Name = data.Name
	c.LegalName = data.LegalName
	c.LoanTypeID = data.LoanTypeID
	c.LoanTypeName = data.LoanTypeName
	c.LoanCodePrefix = data.LoanCodePrefix
	c.InitialCreditLimit = data.InitialCreditLimit
	c.Settings = data.Settings
	c.LoanScheme = data.LoanScheme
	c.EscrowAccount = data.EscrowAccount
	c.CustomSetting = data.CustomSetting
}
