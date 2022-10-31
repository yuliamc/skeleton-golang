package middlewares

import (
	"modalrakyat/skeleton-golang/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// Generic header static API key validator middleware
func (m *MiddlewareAccess) CheckHeaderStaticApiKey(header *string, apiKey *string, errorCode *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		retrievedApiKey := c.Request.Header.Get(*header)
		if retrievedApiKey == "" || retrievedApiKey != *apiKey {
			errors.NewErrorCode(c, *errorCode)
			return
		}

		c.Next()
	}
}
