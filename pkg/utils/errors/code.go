package errors

import (
	"modalrakyat/skeleton-golang/pkg/utils/constant"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Code represent error
const (
	ERROR_MSG_NULL constant.ReserveErrorMessage = iota + 1000
	ERROR_MSG_UNAUTHORIZED_REQUEST
	ERROR_MSG_BAD_REQUEST
	ERROR_MSG_INTERNAL_SERVER_ERROR
	ERROR_MSG_DATA_NOT_FOUND
	ERROR_MSG_INVALID_ENCRYPTION
)

// ERROR_KEYS translate error code to i18n key, determine http status code and error code shown to client
var ERROR_KEYS = map[constant.ReserveErrorMessage]ErrorData{
	ERROR_MSG_NULL:                 *NewErrorData("ERROR_MSG_NULL", http.StatusBadRequest, ERROR_CODE_NULL),
	ERROR_MSG_UNAUTHORIZED_REQUEST: *NewErrorData("ERROR_MSG_UNAUTHORIZED_REQUEST", http.StatusUnauthorized, ERROR_CODE_VALIDATION),
	ERROR_MSG_BAD_REQUEST:          *NewErrorData("ERROR_MSG_BAD_REQUEST", http.StatusBadRequest, ERROR_CODE_VALIDATION),
	ERROR_MSG_DATA_NOT_FOUND:       *NewErrorData("ERROR_MSG_DATA_NOT_FOUND", http.StatusNotFound, ERROR_CODE_GENERAL),
	ERROR_MSG_INVALID_ENCRYPTION:   *NewErrorData("ERROR_MSG_INVALID_ENCRYPTION", http.StatusInternalServerError, ERROR_CODE_GENERAL),
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
	//
	ERROR_CODE_NULL    constant.ReserveErrorCode = 0
	ERROR_CODE_UNKNOWN constant.ReserveErrorCode = 1

	//
	ERROR_CODE_INTERNAL_SERVER_ERROR = 5000

	//
	ERROR_CODE_GENERAL constant.ReserveErrorCode = iota + 1000
	ERROR_CODE_VALIDATION
)

var (
	ResourceNil error = errors.New("create resource nil")
)

func CustomError(message string) error {
	return errors.New(message)
}

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
	errMsgCode := constant.ReserveErrorMessage(se.code)
	return ERROR_KEYS[errMsgCode].ErrorCode
}

func (se *GenericError) GetErrorDataMessageKey() string {
	errMsgCode := constant.ReserveErrorMessage(se.code)
	return ERROR_KEYS[errMsgCode].MessageKey
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
