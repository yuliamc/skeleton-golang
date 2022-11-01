package partner

import (
	"modalrakyat/skeleton-golang/internal/model"
	timeutil "modalrakyat/skeleton-golang/pkg/utils/time"
	"time"
)

func NewPartnerResponse(ID *uint, uniqueID *string, code *string, name *string, createdAt *time.Time, updatedAt *time.Time) PartnerResponse {
	return PartnerResponse{
		ID:        *ID,
		UniqueID:  *uniqueID,
		Code:      *code,
		Name:      *name,
		CreatedAt: timeutil.StrFormat(*createdAt, timeutil.ISO8601TimeWithoutZone),
		UpdatedAt: timeutil.StrFormat(*updatedAt, timeutil.ISO8601TimeWithoutZone),
	}
}

func NewPartnerModelCreatePayload(uniqueID *string, code *string, name *string) model.Partner {
	return model.Partner{
		UniqueID:      *uniqueID,
		Code:          *code,
		Name:          *name,
		LoanTypeID:    99999,
		EscrowAccount: map[string]interface{}{},
	}
}
