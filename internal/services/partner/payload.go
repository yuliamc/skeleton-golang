package partner

type UploadS3Payload struct {
	File   string `json:"file" binding:"required"`
	Folder string `json:"folder"`
}

type CalculatorPayload struct {
	NominalLoan float64 `form:"nominal_loan" binding:"required"`
	Tenor       float64 `form:"tenor" binding:"required"`
}
