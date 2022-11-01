// custom_validator.go
package validator

import (
	"modalrakyat/skeleton-golang/internal/model/enum"
	"modalrakyat/skeleton-golang/pkg/utils/lang"
	"modalrakyat/skeleton-golang/pkg/utils/null"

	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidatorWrapper struct {
	// Validator key
	Key string
	// Return NIL to make it valid
	ValidateField func(value interface{}) interface{}
	// Return validation message
	GetMessage func(e validator.FieldError, vw ValidatorWrapper) string
}

var ValidatorWrappers = []ValidatorWrapper{
	*educationValidator,
	*alphaSpaceValidator,
}

func findValidatorByTag(tag string) *ValidatorWrapper {
	for _, v := range ValidatorWrappers {
		if v.Key == tag {
			return &v
		}
	}
	return nil
}

var educationValidator = &ValidatorWrapper{
	Key: "education_validator",
	ValidateField: func(value interface{}) interface{} {
		if null.IsNil(value) {
			return nil
		}
		education := value.(string)
		if education == "" {
			return nil
		}
		educationPair := enum.GetEducationKeyValuePairs()
		for _, v := range educationPair {
			if v.Value == strings.TrimSpace(education) {
				return nil
			}
		}
		return education
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERROR_MSG_VALIDATION_VALUE", templateData)
		return msg
	},
}

var alphaSpaceValidator = &ValidatorWrapper{
	Key: "alphaspace_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		re, _ := regexp.Compile(`^[a-zA-Z ]+$`)
		if !re.MatchString(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERROR_MSG_VALIDATION_ALPHASPACE", templateData)
		return msg
	},
}

func (o *ValidatorWrapper) Validate(fl validator.FieldLevel, vw ValidatorWrapper) bool {
	return vw.ValidateField(fl.Field().Interface()) == nil
}
