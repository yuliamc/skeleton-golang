package errors

import (
	"modalrakyat/skeleton-golang/pkg/utils/constant"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Code represent error
const (
	NULL int = iota + 1000
	DATA_NOT_FOUND
	CLIENT_AUTH_ERROR
	CLIENT_AUTH_FORBIDDEN
	CLIENT_INFO_NOT_FOUND
	INVALID_ENCRYPTION
	UPLOAD_TO_S3_FAILED
	CREDIT_SCORING_NOT_ELIGIBLE
	UNPROCESSABLE_ERROR
	LOAN_REQUEST_REJECTED_BY_MANUAL_OPERATION
	FAILED_BINDING_PAYLOAD
	INTERNAL_SERVER_ERROR
	UNPROCESSABLE_REQUEST
	SERVICE_DEPENDENCY_ERROR
	DUPLICATE_DATA
	DUPLICATE_REGISTRATION_DATA
	PENDING_REGISTRATION_DATA_OR_ALREADY_LENDER
	DUPLICATE_REFERENCE_ID
	SEND_OTP_FAILED
	VERIFY_OTP_FAILED
	VALIDATION_ERROR
	CANT_CONNECT_BORROWER_NOT_FOUND
	CREDIT_LIMIT_NOT_ENOUGH
	CREDIT_REMAINING_NOT_ENOUGH
	LOAN_SETTING_NOT_FOUND
	LOAN_STILL_EXIST
	SE_BLOCKED
	AUTHENTICATION_FAIL
	TOKEN_UNAUTHORIZED
	INVALID_EMAIL_OR_PASSWORD
	INVALID_OLD_PASSWORD
	USER_REMOVED
	INVALID_RECAPTCHA
	INVALID_RECAPTCHA_DUPLICATE
	UNAUTHORIZED_REQUEST
	PARAM_NOT_VIABLE
	FDC_NOT_ELIGIBLE

	CLIENT_ERROR
	LOAN_REQUESTS_PILOT_APPROVE_FAILED
	LOAN_REQUESTS_BULK_APPROVE_FAILED
	LOAN_REQUESTS_BULK_REJECT_FAILED
	INVALID_LOAN_UPDATE
	LOAN_NOT_EXIST
	INVALID_BATCH_UPDATE
	BORROWER_VERIFICATION_INVALID_STATUS
	INVALID_IP_ADDRESS
	LOAN_ALREADY_PAID
	REPAYMENT_AMOUNT_NOT_ENOUGH
	AGREEMENT_LETTER_NUMBER_EXIST
	INVALID_REPAYMENT_UPDATE
	AGREEMENT_LETTER_IN_PROCESS
	REGISTRATION_BULK_APPROVE_FAILED
	LOAN_CORE_SUBMISSION_FAILED
	REGISTRATION_BULK_REJECT_FAILED
	FORBIDDEN
	FORGOT_PASSWORD_EXCEED_LIMIT
	IMAGE_FIELD_NOT_FOUND
	DUPLICATE_LOAN_REFERENCE_ID
	LOAN_REQUEST_NOT_EXIST
	LOAN_REQUEST_STILL_EXIST
	REGISTRATION_REQUEST_ACQUISITION_NOT_APPLICABLE
	LOAN_REQUEST_ACQUISITION_NOT_APPLICABLE
	CALLBACK_ALREADY_SENT
	CREDIT_EVALUATION_FAILED_FDC
	CREDIT_EVALUATION_FAILED_CREDIT_SCORING
	SALES_EXECUTIVE_BLOCK_FAILED
	SALES_EXECUTIVE_UNBLOCK_FAILED
	PENDING_REGISTRATION_REQUEST_WITH_OTHER_PARTNER
	EXCEEDED_MAX_DATE_RANGE
	CREDIT_CHECKING_CONFIGURATION_NOT_FOUND
	BAD_GATEWAY
	X_SERVERLESS_LOANHUB_API_KEY_INVALID
	X_API_KEY_NOT_VALID
	X_SSO_KEY_NOT_VALID
	INVALID_FAZZCAPITAL_ID
)

// ERROR_KEYS translate error code to i18n key, determine http status code and error code shown to client
var ERROR_KEYS = map[int]ErrorData{
	NULL:                                            *NewErrorData("NULL", http.StatusBadRequest, ERROR_VALIDATION),
	DATA_NOT_FOUND:                                  *NewErrorData("DATA_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	CLIENT_AUTH_ERROR:                               *NewErrorData("CLIENT_AUTH_ERROR", http.StatusUnauthorized, ERROR_VALIDATION),
	CLIENT_AUTH_FORBIDDEN:                           *NewErrorData("CLIENT_AUTH_FORBIDDEN", http.StatusForbidden, ERROR_VALIDATION),
	CLIENT_INFO_NOT_FOUND:                           *NewErrorData("CLIENT_INFO_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_ENCRYPTION:                              *NewErrorData("INVALID_ENCRYPTION", http.StatusBadRequest, ERROR_VALIDATION),
	UPLOAD_TO_S3_FAILED:                             *NewErrorData("UPLOAD_TO_S3_FAILED", http.StatusBadRequest, ERROR_UPLOAD_TO_S3),
	CREDIT_SCORING_NOT_ELIGIBLE:                     *NewErrorData("CREDIT_SCORING_NOT_ELIGIBLE", http.StatusUnprocessableEntity, ERROR_CREDIT_SCORING_REJECTED),
	UNPROCESSABLE_ERROR:                             *NewErrorData("UNPROCESSABLE_ERROR", http.StatusInternalServerError, ERROR_UNKNOWN),
	FAILED_BINDING_PAYLOAD:                          *NewErrorData("FAILED_BINDING_PAYLOAD", http.StatusBadRequest, ERROR_VALIDATION),
	INTERNAL_SERVER_ERROR:                           *NewErrorData("INTERNAL_SERVER_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	UNPROCESSABLE_REQUEST:                           *NewErrorData("UNPROCESSABLE_REQUEST", http.StatusBadRequest, ERROR_VALIDATION),
	SERVICE_DEPENDENCY_ERROR:                        *NewErrorData("SERVICE_DEPENDENCY_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	SEND_OTP_FAILED:                                 *NewErrorData("SEND_OTP_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	VERIFY_OTP_FAILED:                               *NewErrorData("VERIFY_OTP_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	VALIDATION_ERROR:                                *NewErrorData("VALIDATION_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	CANT_CONNECT_BORROWER_NOT_FOUND:                 *NewErrorData("CANT_CONNECT_BORROWER_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_DATA:                                  *NewErrorData("DUPLICATE_DATA", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_REGISTRATION_DATA:                     *NewErrorData("DUPLICATE_REGISTRATION_DATA", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_REFERENCE_ID:                          *NewErrorData("DUPLICATE_REFERENCE_ID", http.StatusBadRequest, ERROR_VALIDATION),
	CREDIT_LIMIT_NOT_ENOUGH:                         *NewErrorData("CREDIT_LIMIT_NOT_ENOUGH", http.StatusBadRequest, ERROR_VALIDATION),
	CREDIT_REMAINING_NOT_ENOUGH:                     *NewErrorData("CREDIT_REMAINING_NOT_ENOUGH", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_SETTING_NOT_FOUND:                          *NewErrorData("LOAN_SETTING_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_STILL_EXIST:                                *NewErrorData("LOAN_STILL_EXIST", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_NOT_EXIST:                                  *NewErrorData("LOAN_NOT_EXIST", http.StatusBadRequest, ERROR_VALIDATION),
	PENDING_REGISTRATION_DATA_OR_ALREADY_LENDER:     *NewErrorData("PENDING_REGISTRATION_DATA_OR_ALREADY_LENDER", http.StatusBadRequest, ERROR_VALIDATION),
	SE_BLOCKED:                                      *NewErrorData("SE_BLOCKED", http.StatusBadRequest, ERROR_SALES_BLOCKED),
	AUTHENTICATION_FAIL:                             *NewErrorData("AUTHENTICATION_FAIL", http.StatusUnauthorized, ERROR_NULL),
	TOKEN_UNAUTHORIZED:                              *NewErrorData("TOKEN_UNAUTHORIZED", http.StatusBadRequest, ERROR_VALIDATION),
	USER_REMOVED:                                    *NewErrorData("USER_REMOVED", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_EMAIL_OR_PASSWORD:                       *NewErrorData("INVALID_EMAIL_OR_PASSWORD", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_OLD_PASSWORD:                            *NewErrorData("INVALID_OLD_PASSWORD", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_RECAPTCHA:                               *NewErrorData("INVALID_RECAPTCHA", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_RECAPTCHA_DUPLICATE:                     *NewErrorData("INVALID_RECAPTCHA_DUPLICATE", http.StatusBadRequest, ERROR_VALIDATION),
	UNAUTHORIZED_REQUEST:                            *NewErrorData("UNAUTHORIZED_REQUEST", http.StatusBadRequest, ERROR_VALIDATION),
	PARAM_NOT_VIABLE:                                *NewErrorData("PARAM_NOT_VIABLE", http.StatusBadRequest, ERROR_VALIDATION),
	CLIENT_ERROR:                                    *NewErrorData("CLIENT_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUESTS_PILOT_APPROVE_FAILED:              *NewErrorData("LOAN_REQUESTS_PILOT_APPROVE_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUESTS_BULK_APPROVE_FAILED:               *NewErrorData("LOAN_REQUESTS_BULK_APPROVE_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUESTS_BULK_REJECT_FAILED:                *NewErrorData("LOAN_REQUESTS_BULK_REJECT_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_LOAN_UPDATE:                             *NewErrorData("INVALID_LOAN_UPDATE", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_BATCH_UPDATE:                            *NewErrorData("INVALID_BATCH_UPDATE", http.StatusBadRequest, ERROR_VALIDATION),
	BORROWER_VERIFICATION_INVALID_STATUS:            *NewErrorData("BORROWER_VERIFICATION_INVALID_STATUS", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_IP_ADDRESS:                              *NewErrorData("INVALID_IP_ADDRESS", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_ALREADY_PAID:                               *NewErrorData("LOAN_ALREADY_PAID", http.StatusBadRequest, ERROR_VALIDATION),
	REPAYMENT_AMOUNT_NOT_ENOUGH:                     *NewErrorData("REPAYMENT_AMOUNT_NOT_ENOUGH", http.StatusBadRequest, ERROR_VALIDATION),
	AGREEMENT_LETTER_NUMBER_EXIST:                   *NewErrorData("AGREEMENT_LETTER_NUMBER_EXIST", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_REPAYMENT_UPDATE:                        *NewErrorData("INVALID_REPAYMENT_UPDATE", http.StatusBadRequest, ERROR_VALIDATION),
	AGREEMENT_LETTER_IN_PROCESS:                     *NewErrorData("AGREEMENT_LETTER_IN_PROCESS", http.StatusBadRequest, ERROR_VALIDATION),
	REGISTRATION_BULK_APPROVE_FAILED:                *NewErrorData("REGISTRATION_BULK_APPROVE_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	REGISTRATION_BULK_REJECT_FAILED:                 *NewErrorData("REGISTRATION_BULK_REJECT_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_CORE_SUBMISSION_FAILED:                     *NewErrorData("LOAN_CORE_SUBMISSION_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	FDC_NOT_ELIGIBLE:                                *NewErrorData("FDC_NOT_ELIGIBLE", http.StatusUnprocessableEntity, ERROR_FDC_REJECTED),
	FORBIDDEN:                                       *NewErrorData("FORBIDDEN", http.StatusForbidden, ERROR_NULL),
	FORGOT_PASSWORD_EXCEED_LIMIT:                    *NewErrorData("FORGOT_PASSWORD_EXCEED_LIMIT", http.StatusBadRequest, ERROR_VALIDATION),
	IMAGE_FIELD_NOT_FOUND:                           *NewErrorData("IMAGE_FIELD_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_LOAN_REFERENCE_ID:                     *NewErrorData("DUPLICATE_LOAN_REFERENCE_ID", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUEST_NOT_EXIST:                          *NewErrorData("LOAN_REQUEST_NOT_EXIST", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUEST_STILL_EXIST:                        *NewErrorData("LOAN_REQUEST_STILL_EXIST", http.StatusBadRequest, ERROR_VALIDATION),
	REGISTRATION_REQUEST_ACQUISITION_NOT_APPLICABLE: *NewErrorData("REGISTRATION_REQUEST_ACQUISITION_NOT_APPLICABLE", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUEST_ACQUISITION_NOT_APPLICABLE:         *NewErrorData("LOAN_REQUEST_ACQUISITION_NOT_APPLICABLE", http.StatusBadRequest, ERROR_VALIDATION),
	CALLBACK_ALREADY_SENT:                           *NewErrorData("CALLBACK_ALREADY_SENT", http.StatusBadRequest, ERROR_VALIDATION),
	CREDIT_EVALUATION_FAILED_FDC:                    *NewErrorData("CREDIT_EVALUATION_FAILED_FDC", http.StatusBadRequest, ERROR_FDC_TIMEOUT),
	CREDIT_EVALUATION_FAILED_CREDIT_SCORING:         *NewErrorData("CREDIT_EVALUATION_FAILED_CREDIT_SCORING", http.StatusBadRequest, ERROR_CREDIT_SCORING_TIMEOUT),
	SALES_EXECUTIVE_BLOCK_FAILED:                    *NewErrorData("SALES_EXECUTIVE_BLOCK_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	SALES_EXECUTIVE_UNBLOCK_FAILED:                  *NewErrorData("SALES_EXECUTIVE_UNBLOCK_FAILED", http.StatusBadRequest, ERROR_VALIDATION),
	LOAN_REQUEST_REJECTED_BY_MANUAL_OPERATION:       *NewErrorData("LOAN_REQUEST_REJECTED_BY_MANUAL_OPERATION", http.StatusBadRequest, ERROR_LOAN_REQUEST_REJECTED_BY_MANUAL_OPERATION),
	EXCEEDED_MAX_DATE_RANGE:                         *NewErrorData("EXCEEDED_MAX_DATE_RANGE", http.StatusBadRequest, ERROR_VALIDATION),
	CREDIT_CHECKING_CONFIGURATION_NOT_FOUND:         *NewErrorData("CREDIT_CHECKING_CONFIGURATION_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	BAD_GATEWAY:                                     *NewErrorData("BAD_GATEWAY", http.StatusBadRequest, ERROR_VALIDATION),
	X_SERVERLESS_LOANHUB_API_KEY_INVALID:            *NewErrorData("X_SERVERLESS_LOANHUB_API_KEY_INVALID", http.StatusUnauthorized, ERROR_VALIDATION),
	X_API_KEY_NOT_VALID:                             *NewErrorData("X_API_KEY_NOT_VALID", http.StatusBadRequest, ERROR_VALIDATION),
	X_SSO_KEY_NOT_VALID:                             *NewErrorData("X_SSO_KEY_NOT_VALID", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_FAZZCAPITAL_ID:                          *NewErrorData("INVALID_FAZZCAPITAL_ID", http.StatusBadRequest, ERROR_VALIDATION),
}

type ErrorData struct {
	// MessageKey is the key that will be used to translate the message. mapping key to id/en.json
	MessageKey string
	// HttpCode is the HTTP status code to be returned to the client.
	HttpCode int
	// ErrorCode is the error code that will be used to identify specific error.
	ErrorCode constant.ReserveErrorCode
}

func NewErrorData(messageKey string, httpCode int, errorCode constant.ReserveErrorCode) *ErrorData {
	return &ErrorData{
		MessageKey: messageKey,
		HttpCode:   httpCode,
		ErrorCode:  errorCode,
	}
}

// ERROR_ is reserved key for error code
const (
	ERROR_NULL    constant.ReserveErrorCode = 0
	ERROR_UNKNOWN constant.ReserveErrorCode = 1
)

// ERROR_ is reserved key for error code
const (
	ERROR_VALIDATION constant.ReserveErrorCode = iota + 4000
	ERROR_FDC_REJECTED
	ERROR_CREDIT_SCORING_REJECTED
	ERROR_UPLOAD_TO_S3
	ERROR_SALES_BLOCKED
	ERROR_FDC_TIMEOUT
	ERROR_CREDIT_SCORING_TIMEOUT
	ERROR_LOAN_REQUEST_REJECTED_BY_MANUAL_OPERATION
)

var (
	ResourceNil error = errors.New("create resource nil")
)

type Err string

func CustomError(message string) error {
	return errors.New(message)
}

// Loanhub error wrapper.
// @TODO Handling thread-safe translation without providing context
type GenericError struct {
	// code that refers to enum iota of "Code represent error"
	code int

	// override default error message
	// message will "ignore/override" translated message from code
	message string

	// errors is the additional error data that shown to client
	details []constant.ErrorDetails

	// callback after the message is translated
	fn func(string) string
}

func (se *GenericError) Error() string {
	return strconv.Itoa(se.GetCode())
}

func (se *GenericError) GetCode() int {
	return se.code
}

func (se *GenericError) GetMessage() string {
	return se.message
}

func (se *GenericError) GetCallback() func(string) string {
	return se.fn
}

func (se *GenericError) GetDetails() []constant.ErrorDetails {
	if se.details == nil {
		return []constant.ErrorDetails{}
	}
	return se.details
}

func (se *GenericError) GetErrorDataCode() constant.ReserveErrorCode {
	return ERROR_KEYS[se.code].ErrorCode
}

func (se *GenericError) GetErrorDataMessageKey() string {
	return ERROR_KEYS[se.code].MessageKey
}

func NewGenericError(code int, options ...func(*GenericError)) error {
	genericError := &GenericError{code: code}

	for _, option := range options {
		option(genericError)
	}

	return genericError
}

func WithCallback(fn func(string) string) func(err *GenericError) {
	return func(err *GenericError) {
		err.fn = fn
	}
}

func OverrideCode(code int) func(err *GenericError) {
	return func(err *GenericError) {
		err.code = code
	}
}

func OverrideErrorMessage(message string) func(err *GenericError) {
	return func(err *GenericError) {
		err.message = message
	}
}

func SetDetails(details []constant.ErrorDetails) func(err *GenericError) {
	return func(err *GenericError) {
		err.details = details
	}
}

// NewGenericErrorWithFn creates a custom error
// Deprecated: use NewGenericError with "OverrideErrorMessage" options instead
// will be removed in the future
func NewGenericErrorWithFn(code int, fn func(string) string) error {
	return &GenericError{code: code, fn: fn}
}
