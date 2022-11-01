package messages

import (
	"fmt"
	"modalrakyat/skeleton-golang/pkg/utils/constant"

	"github.com/gin-gonic/gin"
)

func TranslateCode(ctx *gin.Context, messageCode int) string {
	defaultMessageCode := func(messageCode int) string {
		return fmt.Sprintf("code: %d", messageCode)
	}

	T, ok := ctx.Get("T")
	if ok {
		translator, ok := T.(func(string) (string, error))
		if ok {
			translatedMessage, err := translator(GetKeyByMessageCodeInt(&messageCode))
			if err == nil && len(translatedMessage) > 0 {
				return translatedMessage
			}
		}
	}
	return defaultMessageCode(messageCode)
}

func GetKeyByMessageCodeInt(code *int) string {
	return KEYS[constant.ReserveMessageCode(*code)]
}
