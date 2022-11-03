package store

import (
	hdrs "modalrakyat/skeleton-golang/cmd/apiserver/app/handlers"
	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/internal/middlewares"
	repos "modalrakyat/skeleton-golang/internal/repositories"
	srvs "modalrakyat/skeleton-golang/internal/services"
	"modalrakyat/skeleton-golang/pkg/clients/db"
	"modalrakyat/skeleton-golang/pkg/clients/redis"
)

var (
	// resources
	dbClient    db.DBGormDelegate
	redisClient redis.RedisDelegate

	// hdrs
	BOAuthHandler     *hdrs.BOAuthHandler
	BOAdminHandler    *hdrs.BOAdminHandler
	BOConfigHandler   *hdrs.BOConfigHandler
	MerchantCCHandler *hdrs.MerchantCCHandler
	MerchantVAHandler *hdrs.MerchantVAHandler
	PartnerHandler    *hdrs.PartnerHandler
	CallbackHandler   *hdrs.CallbackHandler

	// srvs
	partnerService srvs.PartnerService

	// repos
	partnerRepo repos.PartnerRepo
	txRepo      repos.TxRepo

	// middlewares
	AccessMiddleware         middlewares.AccessMiddleware
	BackofficeAuthMiddleware middlewares.BackofficeAuthMiddleware
)

// Init application global variable with single instance
func InitDI() {
	// setup resources
	dbClient = db.NewDBdelegate(config.Config.DB.Debug)
	dbClient.Init()

	redisClient = redis.NewRedisDel()
	redisClient.Init()

	// setup components
	// repos
	partnerRepo = repos.NewPartnerRepo(dbClient)
	txRepo = repos.NewTxRepo(dbClient)

	// services
	partnerService = srvs.NewPartnerService(partnerRepo, txRepo)

	// hdrs
	PartnerHandler = hdrs.NewPartnerHandler(partnerService)
	BOAdminHandler = hdrs.NewBOAdminHandler()
	BOConfigHandler = hdrs.NewBOConfigHandler(redisClient)

	// middleware
	AccessMiddleware = middlewares.NewAccessMiddleware(redisClient, partnerService)
	BackofficeAuthMiddleware = middlewares.NewBackofficeAuthMiddleware()
}
