package partner

import (
	"context"
	repos "modalrakyat/skeleton-golang/internal/repositories"
)

type PartnerService interface {
	FindByID(ctx context.Context, ID *uint) (*PartnerResponse, error)
	FindAll(ctx context.Context) (*[]PartnerResponse, *int64, error)
	Create(ctx context.Context, payload *CreatePartnerPayload) error
}

type partnerService struct {
	partnerRepo repos.PartnerRepo
	txRepo      repos.TxRepo
}

func NewPartnerService(partnerRepo repos.PartnerRepo, txRepo repos.TxRepo) PartnerService {
	return &partnerService{
		partnerRepo: partnerRepo,
		txRepo:      txRepo,
	}
}
