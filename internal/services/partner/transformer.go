package partner

import (
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
