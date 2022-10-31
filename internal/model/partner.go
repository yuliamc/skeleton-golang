package model

import (
	"time"
)

// TableName overrides the table name used by Partner to `partner`
func (Partner) TableName() string {
	return "partner"
}

type Partner struct {
	ID                 uint      `gorm:"primarykey"`
	UniqueID           string    `gorm:"column:unique_id;type:uuid"`
	Code               string    `gorm:"column:code"`
	Name               string    `gorm:"column:name"`
	LegalName          string    `gorm:"column:legal_name"`
	LoanTypeID         uint      `gorm:"column:loan_type_id"`
	LoanTypeName       string    `gorm:"column:loan_type_name"`
	LoanCodePrefix     string    `gorm:"column:loan_code_prefix"`
	InitialCreditLimit float64   `gorm:"column:initial_credit_limit"`
	Settings           JsonArr   `gorm:"column:settings;type:text"`
	LoanScheme         JsonArr   `gorm:"column:loan_scheme;type"`
	EscrowAccount      Json      `gorm:"column:escrow_account;type:json"`
	CustomSetting      Json      `gorm:"column:custom_setting;type:json"`
	CreatedAt          time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:datetime"`
}
