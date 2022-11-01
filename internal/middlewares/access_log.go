package middlewares

import (
	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/logs"
	"modalrakyat/skeleton-golang/pkg/utils/net"
	stringer "modalrakyat/skeleton-golang/pkg/utils/strings"
	timeutil "modalrakyat/skeleton-golang/pkg/utils/time"

	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := timeutil.Now()

		requestParser := net.HTTPRequestParser{}
		requestParser.ParseRequestBody(c)

		c.Next()

		end := timeutil.Now()
		latency := end.Sub(t)

		var headers map[string][]string = c.Request.Header
		if c.Request.Header.Get("Authorization") != "" {
			headers["Authorization"][0] = stringer.MaskUUIDV4(c.Request.Header.Get("Authorization"))
		}

		fields := logs.Fields{
			"client_ip":       net.GetClientIpAddress(c),
			"client_os":       c.Request.Header.Get("Client-OS"),
			"client_version":  c.Request.Header.Get("Client-Version"),
			"request_id":      c.GetString("RequestId"),
			"request_uri":     c.Request.RequestURI,
			"method":          c.Request.Method,
			"user_agent":      c.Request.UserAgent(),
			"referer":         c.Request.Referer(),
			"mode":            config.Config.System.Mode,
			"host":            c.Request.Host,
			"path":            c.Request.URL.Path,
			"params":          c.Request.URL.RawQuery,
			"lang":            c.Request.Header.Get("Accept-Language"),
			"status":          c.Writer.Status(),
			"process_time":    latency.String(),
			"process_time_ns": latency.Nanoseconds(),
			"request_body":    requestParser.GetRequestBody(c),
			"request_header":  c.Request.Header,
			"type_str":        "GIN",
		}

		cl := logs.Log.WithFields(fields)
		routePathParamMap := make(map[string]interface{})
		for _, p := range c.Params {
			routePathParamMap[p.Key] = p.Value
		}
		cl = cl.WithFields(logs.Fields{
			"route_path_params": routePathParamMap,
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, v := range c.Errors {
				cl = cl.WithFields(logs.Fields{
					"error_string": errors.ToString(v.Err),
					"error_stack":  errors.GetStack(v.Err),
				})
				break
			}
			logs.Log.Errorln(cl)
		} else {
			logs.Log.Infoln(cl)
		}
	}
}
