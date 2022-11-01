package model

import (
	"time"
)

// TableName overrides the table name used by Partner to `partner`
func (Partner) TableName() string {
	return "partner"
}

type Partner struct {
	ID            uint      `gorm:"primarykey"`
	UniqueID      string    `gorm:"column:unique_id;type:uuid"`
	Code          string    `gorm:"column:code"`
	Name          string    `gorm:"column:name"`
	LoanTypeID    uint      `gorm:"column:loan_type_id"`
	EscrowAccount Json      `gorm:"column:escrow_account;type:json"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime"`
}
