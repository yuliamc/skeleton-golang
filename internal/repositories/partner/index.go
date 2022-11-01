package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/clients/db"
)

type PartnerRepo interface {
	Find(ctx context.Context, where *model.Partner) (*model.Partner, error)
}

type partnerRepo struct {
	dbClient db.DBGormDelegate
}

func NewPartnerRepo(dbClient db.DBGormDelegate) PartnerRepo {
	return &partnerRepo{
		dbClient: dbClient,
	}
}
