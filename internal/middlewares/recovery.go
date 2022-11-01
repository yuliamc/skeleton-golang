package middlewares

import (
	"net/http"

	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/logs"
	"modalrakyat/skeleton-golang/pkg/utils/net"
	netutil "modalrakyat/skeleton-golang/pkg/utils/net"
	stringutil "modalrakyat/skeleton-golang/pkg/utils/strings"
	timeutil "modalrakyat/skeleton-golang/pkg/utils/time"

	"github.com/gin-gonic/gin"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// Recovery returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func Recovery(mode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := timeutil.Now()

		requestParser := net.HTTPRequestParser{}
		requestParser.ParseRequestBody(c)

		defer func() {
			if err := recover(); err == nil {
				return
			}

			err := recover()
			end := timeutil.Now()
			latency := end.Sub(t)

			var headers map[string][]string = c.Request.Header
			if c.Request.Header.Get("Authorization") != "" {
				headers["Authorization"][0] = stringutil.MaskUUIDV4(c.Request.Header.Get("Authorization"))
			}

			fields := logs.Fields{
				"client_ip":       netutil.GetClientIpAddress(c),
				"client_os":       c.Request.Header.Get("Client-OS"),
				"client_version":  c.Request.Header.Get("Client-Version"),
				"request_id":      c.GetString("RequestId"),
				"request_uri":     c.Request.RequestURI,
				"method":          c.Request.Method,
				"handler":         c.HandlerName(),
				"user_agent":      c.Request.UserAgent(),
				"referer":         c.Request.Referer(),
				"mode":            config.Config.System.Mode,
				"host":            c.Request.Host,
				"path":            c.Request.URL.Path,
				"params":          c.Request.URL.RawQuery,
				"lang":            c.Request.Header.Get("Accept-Language"),
				"status":          http.StatusInternalServerError,
				"process_time":    latency.String(),
				"process_time_ns": latency.Nanoseconds(),
				"error_string":    errors.ToString(err),
				"error_stack":     errors.GetStack(err),
				"request_body":    requestParser.GetRequestBody(c),
				"request_header":  c.Request.Header,
				"type_str":        "ERR-PANIC",
			}

			cl := logs.Log.WithFields(fields)
			routePathParamMap := make(map[string]interface{})
			for _, p := range c.Params {
				routePathParamMap[p.Key] = p.Value
			}
			cl = cl.WithFields(logs.Fields{
				"route_path_params": routePathParamMap,
			})

			c.AbortWithStatusJSON(500, api.Error{
				Message: errors.Translate(c, int(errors.ERROR_MSG_INTERNAL_SERVER_ERROR)),
			})
		}()
		c.Next()
	}
}
