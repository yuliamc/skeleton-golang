package middlewares

import (
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *BackofficeAuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gets Authorization header.
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			errors.ResponseErrorWithCode(c, int(errors.ERROR_MSG_UNAUTHORIZED_REQUEST))
			return
		}

		// Splits the Authorization value.
		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			if token[0] != "Bearer" {
				errors.ResponseErrorWithCode(c, int(errors.ERROR_MSG_UNAUTHORIZED_REQUEST))
				return
			} else {
				errors.ResponseErrorWithCode(c, int(errors.ERROR_MSG_UNAUTHORIZED_REQUEST))
				return
			}
		}

		// Converts sid into admin object by searching it on redis,
		// then stores it on session so we can now who uses the endpoint.
		// SID := token[1]
		// var session *admin_auth.MCAdminSessionResponse
		// session, err := m.adminAuthService.GetAdminSession(c, SID)
		// if err != nil {
		// 	errors.ErrorCode(c, http.StatusUnauthorized, errors.CLIENT_AUTH_ERROR)
		// 	return
		// } else {
		// 	c.Set("admin", *session)
		// }

		c.Next()
	}
}
