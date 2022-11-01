package constant

type Creator struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PartnerID   uint   `json:"partner_id"`
	Email       string `json:"email"`
	UniqueID    string `json:"unique_id"`
	PhoneNumber string `json:"phone_number"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ReserveErrorCode int
type ReserveErrorMessage int
type ReserveMessageCode int

type ErrorDetails struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
