package middlewares

import (
	srvs "modalrakyat/skeleton-golang/internal/services"
	"modalrakyat/skeleton-golang/pkg/clients/redis"
)

type AccessMiddleware struct {
	redisClient    redis.RedisDelegate
	partnerService srvs.PartnerService
}

func NewAccessMiddleware(redisClient redis.RedisDelegate, partnerService srvs.PartnerService) AccessMiddleware {
	return AccessMiddleware{
		redisClient:    redisClient,
		partnerService: partnerService,
	}
}

type BackofficeAuthMiddleware struct {
}

func NewBackofficeAuthMiddleware() BackofficeAuthMiddleware {
	return BackofficeAuthMiddleware{}
}
