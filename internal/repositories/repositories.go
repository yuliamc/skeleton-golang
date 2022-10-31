package repos

import (
	"modalrakyat/skeleton-golang/internal/repositories/partner"
	"modalrakyat/skeleton-golang/internal/repositories/tx"
)

// put repos alias
type (
	PartnerRepo = partner.PartnerRepo
	TxRepo      = tx.TxRepo
)

var (
	NewPartnerRepo = partner.NewPartnerRepo
	NewTxRepo      = tx.NewTxRepo
)
