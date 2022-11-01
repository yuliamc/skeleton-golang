package partner

func NewPartnerResponse(ID *uint, uniqueID *string, code *string, name *string) PartnerResponse {
	return PartnerResponse{
		ID:       *ID,
		UniqueID: *uniqueID,
		Code:     *code,
		Name:     *name,
	}
}
