package middlewares

import (
	"modalrakyat/skeleton-golang/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// Generic header static API key validator middleware
func StaticApiKey(apiKey *string) gin.HandlerFunc {
	header := "x-api-key"
	return func(c *gin.Context) {
		retrievedApiKey := c.Request.Header.Get(header)
		if retrievedApiKey == "" || retrievedApiKey != *apiKey {
			errors.ResponseErrorWithCode(c, int(errors.ERROR_MSG_UNAUTHORIZED_REQUEST))
			return
		}

		c.Next()
	}
}
