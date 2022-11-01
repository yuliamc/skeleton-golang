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
	PartnerHandler *hdrs.PartnerHandler

	// srvs
	partnerService srvs.PartnerService

	// repos
	partnerRepo repos.PartnerRepo
	txRepo      repos.TxRepo

	//middleware
	accessMiddleware middlewares.MiddlewareAccess
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
	partnerService = srvs.NewPartnerService(&partnerRepo)

	// hdrs
	PartnerHandler = hdrs.NewPartnerHandler(&partnerService)

	// middleware
	accessMiddleware = middlewares.NewMiddlewareAccess(&redisClient, &partnerService)

}
