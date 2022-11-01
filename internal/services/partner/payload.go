package partner

type Uri struct {
	PartnerID uint `uri:"id" binding:"required"`
}
