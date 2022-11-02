package hds

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/bo_admin"
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/bo_auth"
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/callback"
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/merchant_cc"
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/merchant_va"
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/partner"
)

// put handlers alias
type (
	BOAuthHandler     = bo_auth.BOAuthHandler
	BOAdminHandler    = bo_admin.BOAdminHandler
	MerchantCCHandler = merchant_cc.MerchantCCHandler
	MerchantVAHandler = merchant_va.MerchantVAHandler
	CallbackHandler   = callback.CallbackHandler
	PartnerHandler    = partner.PartnerHandler
)

var (
	NewBOAuthHandler     = bo_auth.NewBOAuthHandler
	NewBOAdminHandler    = bo_admin.NewBOAdminHandler
	NewMerchantCCHandler = merchant_cc.NewMerchantCCHandler
	NewMerchantVAHandler = merchant_va.NewMerchantVAHandler
	NewCallbackHandler   = callback.NewCallbackHandler
	NewPartnerHandler    = partner.NewPartnerHandler
)
