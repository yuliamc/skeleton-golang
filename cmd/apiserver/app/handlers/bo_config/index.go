package bo_config

import "modalrakyat/skeleton-golang/pkg/clients/redis"

type BOConfigHandler struct {
	redisClient redis.RedisDelegate
}

func NewBOConfigHandler(redisClient redis.RedisDelegate) *BOConfigHandler {
	return &BOConfigHandler{
		redisClient: redisClient,
	}
}
