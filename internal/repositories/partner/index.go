package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/clients/db"
)

type PartnerRepo interface {
	Find(ctx context.Context, where *model.Partner) (*model.Partner, error)
	FindAll(ctx context.Context) (*[]model.Partner, error)
	CountAll(ctx context.Context) (*int64, error)
	Create(ctx context.Context, partner *model.Partner) error
}

type partnerRepo struct {
	dbClient db.DBGormDelegate
}

func NewPartnerRepo(dbClient db.DBGormDelegate) PartnerRepo {
	return &partnerRepo{
		dbClient: dbClient,
	}
}
