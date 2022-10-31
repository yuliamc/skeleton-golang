package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/clients/db"
)

type PartnerRepo interface {
	// Common / Apps
	GetByID(ctx context.Context, ID uint) (*model.Partner, error)

	// Microlending
	MLGetDropdown(ctx context.Context, fields ...string) ([]*model.Partner, error)
}

type partnerRepo struct {
	dbdget db.DBGormDelegate
}

func NewPartnerRepo(dbdget db.DBGormDelegate) PartnerRepo {
	return &partnerRepo{
		dbdget: dbdget,
	}
}
