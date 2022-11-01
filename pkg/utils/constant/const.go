package constant

//todo: harus di satuin creator sama accessor, jadiin global
type Creator struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PartnerID   uint   `json:"partner_id"`
	Email       string `json:"email"`
	UniqueID    string `json:"unique_id"`
	PhoneNumber string `json:"phone_number"`
}

//todo: harus di satuin creator sama accessor, jadiin global
type Accessor struct {
	ID        uint
	PartnerID uint
}

// Key value pair struct
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ReserveErrorCode int
type ReserveErrorMessage int

type ErrorDetails struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
