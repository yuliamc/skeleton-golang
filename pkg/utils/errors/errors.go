package errors

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/constant"
)

// E messageCode refers to errors/code.go enum iota of "Code represent error"
// Deprecated: use NewE instead
func E(ctx *gin.Context, httpCode int, messageCode int, err error) {
	ctx.Status(httpCode)
	if err != nil {
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: err})
	}

	// Overrides the error from the service will be displayed on response.
	genericError, genericErrorOk := err.(*GenericError)
	if !genericErrorOk {
		ctx.JSON(http.StatusInternalServerError, api.Error{
			Message: "Internal Server Error",
			Code:    ERROR_UNKNOWN,
			Details: []constant.ErrorDetails{
				{Key: "error", Value: err.Error()},
			},
		})
	} else {
		messageCode = genericError.GetCode()
		messageKey := ERROR_KEYS[messageCode].MessageKey
		errorCode := ERROR_KEYS[messageCode].ErrorCode

		T, ok := ctx.Get("T")
		if ok {
			translator, ok := T.(func(string) (string, error))

			if ok {
				translatedMessage, err := translator(messageKey)
				if err == nil && len(translatedMessage) > 0 {
					if genericErrorOk && genericError.fn != nil {
						// Transforms the message for custom message.
						translatedMessage = genericError.fn(translatedMessage)
					}

					ctx.JSON(httpCode, api.Error{Message: translatedMessage, Code: errorCode, Details: genericError.GetDetails()})
				} else {
					if genericErrorOk && genericError.fn != nil {
						// Transforms the message for custom message.
						translatedMessage = genericError.fn(translatedMessage)
						ctx.JSON(httpCode, api.Error{Message: translatedMessage, Code: errorCode, Details: genericError.GetDetails()})
					} else {
						ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
					}
				}
			} else {
				ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
			}
		} else {
			ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
		}
	}

	ctx.Abort()
}

// ErrorCode messageCode refers to errors/code.go enum iota of "Code represent error"
// Deprecated: use NewErrorCode instead
func ErrorCode(ctx *gin.Context, httpCode int, messageCode int) {
	ctx.Status(httpCode)

	messageKey := ERROR_KEYS[messageCode].MessageKey
	errorCode := ERROR_KEYS[messageCode].ErrorCode

	T, ok := ctx.Get("T")
	if ok {
		translator, ok := T.(func(string) (string, error))
		if ok {
			translatedMessage, err := translator(messageKey)
			ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(translatedMessage)})
			if err == nil && len(translatedMessage) > 0 {
				ctx.JSON(httpCode, api.Error{Message: translatedMessage, Code: errorCode, Details: []constant.ErrorDetails{}})
			} else {
				ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: []constant.ErrorDetails{}})
			}
		} else {
			ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: []constant.ErrorDetails{}})
		}
	} else {
		ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: []constant.ErrorDetails{}})
	}

	ctx.Abort()
}

// ErrorString construct a new error with message
// Deprecated: use NewErrorString instead
func ErrorString(ctx *gin.Context, httpCode int, message string) {
	ctx.Status(httpCode)

	T, ok := ctx.Get("T")
	if ok {
		translator, ok := T.(func(string) (string, error))
		if ok {
			translatedMessage, err := translator(message)
			if err == nil {
				ctx.JSON(httpCode, api.Error{Message: translatedMessage, Code: ERROR_VALIDATION, Details: []constant.ErrorDetails{}})
				ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(translatedMessage)})
			} else {
				ctx.JSON(httpCode, api.Error{Message: message, Code: ERROR_VALIDATION, Details: []constant.ErrorDetails{}})
				ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(message)})
			}
		} else {
			ctx.JSON(httpCode, api.Error{Message: message, Code: ERROR_VALIDATION, Details: []constant.ErrorDetails{}})
			ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(message)})
		}
	} else {
		ctx.JSON(httpCode, api.Error{Message: message, Code: ERROR_VALIDATION, Details: []constant.ErrorDetails{}})
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(message)})
	}

	ctx.Abort()
}

// NewE construct an error (usually the error is GenericError).
func NewE(ctx *gin.Context, err error) {
	if err != nil {
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: err})
	}

	genericError, genericErrorOk := err.(*GenericError)
	if !genericErrorOk {
		ctx.JSON(http.StatusInternalServerError, api.Error{
			Message: "Internal Server Error",
			Code:    ERROR_UNKNOWN,
			Details: []constant.ErrorDetails{
				{Key: "error", Value: err.Error()},
			},
		})
	} else {
		messageCode := genericError.GetCode()
		errorData, errorDataOk := ERROR_KEYS[messageCode]
		if !errorDataOk {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{
				Message: "Internal Server Error",
				Code:    ERROR_UNKNOWN,
				Details: []constant.ErrorDetails{
					{Key: "error", Value: genericError.Error()},
				},
			})
			return
		}

		var (
			messageKey = errorData.MessageKey
			httpCode   = errorData.HttpCode
			errorCode  = errorData.ErrorCode
		)

		// override error message
		if genericError.GetMessage() != "" && genericError.GetCallback() == nil {
			ctx.JSON(httpCode, api.Error{Message: genericError.GetMessage(), Code: errorCode, Details: genericError.GetDetails()})
		} else {
			if T, tOk := ctx.Get("T"); tOk {
				if translator, translatorOk := T.(func(string) (string, error)); translatorOk {
					if translatedMessage, err := translator(messageKey); err == nil && len(translatedMessage) > 0 {
						if fn := genericError.GetCallback(); fn != nil {
							translatedMessage = fn(translatedMessage)
						}
						ctx.JSON(httpCode, api.Error{Message: translatedMessage, Code: errorCode, Details: genericError.GetDetails()})
					} else {
						ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
					}
				} else {
					ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
				}
			} else {
				ctx.JSON(httpCode, api.Error{Message: fmt.Sprintf("error code: %s", strconv.Itoa(messageCode)), Code: errorCode, Details: genericError.GetDetails()})
			}
		}
	}
	ctx.Abort()
}

// NewErrorString construct a new error with message
func NewErrorString(ctx *gin.Context, message string, options ...func(*GenericError)) {
	errorOptions := []func(*GenericError){OverrideErrorMessage(message)}
	errorOptions = append(errorOptions, options...)
	NewE(ctx, NewGenericError(NULL, errorOptions...))
}

// NewErrorCode construct a new error with messageCode.
// messageCode refers to errors/code.go enum iota of "Code represent error"
func NewErrorCode(ctx *gin.Context, messageCode int, options ...func(*GenericError)) {
	NewE(ctx, NewGenericError(messageCode, options...))
}

func Translate(ctx context.Context, messageCode int) string {
	context := ctx.(*gin.Context)
	T, ok := context.Get("T")
	if ok {
		translator, ok := T.(func(string) (string, error))
		if ok {
			messageKey := ERROR_KEYS[messageCode].MessageKey
			translatedMessage, err := translator(messageKey)
			if err == nil && len(translatedMessage) > 0 {
				return translatedMessage
			}
		}
	}
	message := fmt.Sprint(messageCode)
	return message
}

func ToString(err interface{}) string {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			return se.Error()
		} else {
			return ne.Error()
		}
	}
	if genericError, ok := err.(*GenericError); ok {
		return genericError.GetErrorDataMessageKey()
	}
	if e, ok := (err).(error); ok {
		return e.Error()
	}
	return fmt.Sprint(err)
}

func GetStack(err interface{}) string {
	const MAX_STACK_LINE = 10
	var countStackLine int = 0
	var stringBuffer bytes.Buffer

	stack := fmt.Sprintf("%s", debug.Stack())
	stacks := strings.Split(stack, "\n")
	for _, v := range stacks {
		if countStackLine >= MAX_STACK_LINE {
			break
		}

		start := strings.Index(v, "/loanhub-service")
		if start < 0 {
			continue
		}
		end := len(v)
		newText := strings.TrimSpace(v[start:end])
		stringBuffer.WriteString(newText)
		stringBuffer.WriteString("\n")
		countStackLine++
	}
	return stringBuffer.String()
}
