package partner

type Uri struct {
	PartnerID uint `uri:"id" binding:"required"`
}

type CreatePartnerPayload struct {
	Code string `json:"code" binding:"required" validate:"min=8,max=50"`
	Name string `json:"name" binding:"required" validate:"min=3,max=100"`
}
