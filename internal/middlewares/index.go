package middlewares

import (
	srvs "modalrakyat/skeleton-golang/internal/services"
	"modalrakyat/skeleton-golang/pkg/clients/redis"
)

type MiddlewareAccess struct {
	redisDel       redis.RedisDelegate
	partnerService srvs.PartnerService
}

func NewMiddlewareAccess(redisDel redis.RedisDelegate, partnerService srvs.PartnerService) MiddlewareAccess {
	return MiddlewareAccess{
		redisDel:       redisDel,
		partnerService: partnerService,
	}
}
