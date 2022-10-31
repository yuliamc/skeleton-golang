// middleware / lang.go
package middlewares

import (
	"modalrakyat/skeleton-golang/pkg/utils/lang"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func LanguageAccept() gin.HandlerFunc {

	// Validation Accept-Language
	return func(c *gin.Context) {
		langCode := c.GetHeader("Accept-Language")
		switch langCode {
		case "", language.Indonesian.String():
			c.Set("T", lang.GetTtranlateFunc(langCode))
			lang.CurrentTranslation.SetTranslation(lang.GetMappingFunc(langCode))
		default:
			c.Set("T", lang.GetTtranlateFunc(
				language.AmericanEnglish.String(),
			))
			lang.CurrentTranslation.SetTranslation(lang.GetMappingFunc(language.AmericanEnglish.String()))
		}

		c.Next()
	}
}
