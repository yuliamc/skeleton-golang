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
	// hdrs
	LoanSchemeHandler *hdrs.LoanSchemeHandler

	// srvs
	PartnerService srvs.PartnerService

	// repos
	PartnerRepo repos.PartnerRepo
	TxRepo      repos.TxRepo

	//middleware
	MiddlewareAccess middlewares.MiddlewareAccess
)

// Init application global variable with single instance
func InitDI() {
	// setup resources
	dbdget := db.NewDBdelegate(config.Config.DB.Debug)
	dbdget.Init()
	redisDel := redis.NewRedisDel()
	redisDel.Init()

	// setup components
	// repos
	PartnerRepo = repos.NewPartnerRepo(dbdget)
	TxRepo = repos.NewTxRepo(dbdget)

	// services
	PartnerService = srvs.NewPartnerService(PartnerRepo)

	// hdrs
	LoanSchemeHandler = hdrs.NewLoanSchemeHandler(RegistrationPartnerService, LoanSchemeService)

	// middleware
	MiddlewareAccess = middlewares.NewMiddlewareAccess(redisDel, PartnerService)

}
