package partner

type PartnerResponse struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	UniqueID string `json:"unique_id"`
	Name     string `json:"name"`
}
