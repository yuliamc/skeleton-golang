package net

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type HTTPRequestParser struct {
	reqBody map[string]interface{}
}

func (s *HTTPRequestParser) ParseRequestBody(c *gin.Context) {
	ByteBody, _ := c.GetRawData()
	if len(ByteBody) != 0 {
		err := json.Unmarshal(ByteBody, &s.reqBody)
		if err != nil {
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
	}
}

func (s *HTTPRequestParser) GetRequestBody(c *gin.Context) map[string]interface{} {
	c.ShouldBind(s.reqBody)

	return s.reqBody
}
