package partner

import (
	srvs "modalrakyat/skeleton-golang/internal/services"
)

type PartnerHandler struct {
	PartnerService srvs.PartnerService
}

func NewPartnerHandler(partnerService srvs.PartnerService) *PartnerHandler {
	return &PartnerHandler{
		PartnerService: partnerService,
	}
}
