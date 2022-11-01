package middlewares

import (
	srvs "modalrakyat/skeleton-golang/internal/services"
	"modalrakyat/skeleton-golang/pkg/clients/redis"
)

type MiddlewareAccess struct {
	redisClient    redis.RedisDelegate
	partnerService srvs.PartnerService
}

func NewMiddlewareAccess(redisClient *redis.RedisDelegate, partnerService *srvs.PartnerService) MiddlewareAccess {
	return MiddlewareAccess{
		redisClient:    *redisClient,
		partnerService: *partnerService,
	}
}
