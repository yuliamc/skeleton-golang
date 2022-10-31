package hds

import (
	"modalrakyat/skeleton-golang/cmd/apiserver/app/handlers/loan_scheme"
)

// put handlers alias
type (
	LoanSchemeHandler = loan_scheme.LoanSchemeHandler
)

var (
	NewLoanSchemeHandler = loan_scheme.NewLoanSchemeHandler
)
