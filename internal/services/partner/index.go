package partner

import (
	"context"
	repos "modalrakyat/skeleton-golang/internal/repositories"
	"modalrakyat/skeleton-golang/pkg/clients/redis"
)

type PartnerService interface {
	// Common / Apps
	GetByID(ctx context.Context, ID uint) (*PartnerResponse, error)
}

type partnerService struct {
	partnerRepo repos.PartnerRepo
	redisDel    redis.RedisDelegate
}

func NewPartnerService(partnerRepo repos.PartnerRepo, redisDel redis.RedisDelegate) PartnerService {
	return &partnerService{
		partnerRepo: partnerRepo,
		redisDel:    redisDel,
	}
}
