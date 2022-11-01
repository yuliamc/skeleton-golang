package partner

type Uri struct {
	PartnerID uint `uri:"id" binding:"required"`
}

type CreatePartnerPayload struct {
	Code string `json:"code" binding:"required" validate:"min=3,max=10"`
	Name string `json:"name" binding:"required" validate:"min=3,max=100"`
}
