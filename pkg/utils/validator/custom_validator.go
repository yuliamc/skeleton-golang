// custom_validator.go
package validator

import (
	"math"
	"modalrakyat/skeleton-golang/internal/model/enum"
	"modalrakyat/skeleton-golang/pkg/utils/lang"
	"modalrakyat/skeleton-golang/pkg/utils/null"
	timeutil "modalrakyat/skeleton-golang/pkg/utils/time"

	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	*profilingDataValidator,
	*jobValidator,
	*educationValidator,
	*maritalStatusValidator,
	*religionValidator,
	*genderValidator,
	*whitelistIpValidator,
	*multipleOfValidator,
	*alphaSpaceValidator,
	*dateYYYYMMDDValidator,
	*callbackEventTypeValidator,
	*registrationStatusValidator,
	*emergencyContactRelationshipValidator,
	*registrationImageTypeValidator,
	*seTypeValidator,
	*passwordValidator,
	*latitudeValidator,
	*longitudeValidator,
	*nameValidator,
}

func findValidatorByTag(tag string) *ValidatorWrapper {
	for _, v := range ValidatorWrappers {
		if v.Key == tag {
			return &v
		}
	}
	return nil
}

var jobValidator = &ValidatorWrapper{
	Key: "job_validator",
	ValidateField: func(value interface{}) interface{} {
		if null.IsNil(value) {
			return nil
		}
		job := value.(string)
		if job == "" {
			return nil
		}
		jobPair := enum.GetJobKeyValuePairs()
		for _, v := range jobPair {
			if v.Value == strings.TrimSpace(job) {
				return nil
			}
		}
		return job
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_JOB", templateData)
		return msg
	},
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
		msg, _ := lang.CurrentTranslation.Translate("ERR_EDUCATION", templateData)
		return msg
	},
}

var maritalStatusValidator = &ValidatorWrapper{
	Key: "marital_status_validator",
	ValidateField: func(value interface{}) interface{} {
		if null.IsNil(value) {
			return nil
		}
		maritalStatus := value.(string)
		if maritalStatus == "" {
			return nil
		}
		maritalStatusPair := enum.GetMaritalStatusKeyValuePairs()
		for _, v := range maritalStatusPair {
			if v.Value == strings.TrimSpace(maritalStatus) {
				return nil
			}
		}
		return maritalStatus
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_MARITAL_STATUS", templateData)
		return msg
	},
}

var religionValidator = &ValidatorWrapper{
	Key: "religion_validator",
	ValidateField: func(value interface{}) interface{} {
		if null.IsNil(value) {
			return nil
		}
		religion := value.(string)
		if religion == "" {
			return nil
		}
		religionPair := enum.GetReligionKeyValuePairs()
		for _, v := range religionPair {
			if v.Value == strings.TrimSpace(religion) {
				return nil
			}
		}
		return religion
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_RELIGION", templateData)
		return msg
	},
}

var genderValidator = &ValidatorWrapper{
	Key: "gender_validator",
	ValidateField: func(value interface{}) interface{} {
		if null.IsNil(value) {
			return nil
		}
		gender := value.(string)
		if gender == "" {
			return nil
		}
		genderPair := enum.GetGenderKeyValuePairs()
		for _, v := range genderPair {
			if v.Key == strings.TrimSpace(gender) {
				return nil
			}
		}
		return gender
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_GENDER", templateData)
		return msg
	},
}

var profilingDataValidator = &ValidatorWrapper{
	Key: "profiling_data_validator",
	ValidateField: func(value interface{}) interface{} {
		profilingData := value.(map[string]interface{})
		isNil := func(v interface{}) bool {
			return v == nil
		}
		isString := func(v interface{}) bool {
			if isNil(v) {
				return false
			}
			rv := reflect.TypeOf(v)
			return rv.Kind() == reflect.String
		}
		isStringNotEmpty := func(v interface{}) bool {
			if !isString(v) {
				return false
			}
			return len(strings.TrimSpace(v.(string))) > 0
		}
		isNumeric := func(v interface{}) bool {
			if isNil(v) {
				return false
			}
			if str, ok := v.(string); ok {
				re, _ := regexp.Compile(`^[0-9]+$`)
				if re.MatchString(str) {
					return true
				}
				return false
			}
			if _, okk := v.(float64); okk {
				return okk
			}
			return false
		}
		isLatitude := func(v interface{}) bool {
			return latitudeValidator.ValidateField(v) == nil
		}
		isLongitude := func(v interface{}) bool {
			return longitudeValidator.ValidateField(v) == nil
		}
		isMap := func(v interface{}) bool {
			if isNil(v) {
				return false
			}
			return reflect.TypeOf(v).Kind() == reflect.Map
		}
		isSalesExecutiveValid := func(v interface{}) bool {
			if isNil(v) {
				return true
			}
			se := v.(map[string]interface{})
			// LLH-242, Support backward compatibility, se_code can be null or empty
			if len(se["se_code"].(string)) > 40 || len(se["se_name"].(string)) > 40 ||
				len(se["hoa_name"].(string)) > 255 || len(se["hor_name"].(string)) > 25 {
				return false
			}
			return isMap(v)
		}
		isDate := func(v interface{}) bool {
			if isNil(v) {
				return false
			}
			_, err := time.Parse(timeutil.ISO8601TimeWithoutZone, v.(string))
			return err == nil
		}
		isUserLevel := func(v interface{}) bool {
			result := seTypeValidator.ValidateField(v)
			return result == nil
		}
		isLoanLevel := func(v interface{}) bool {
			if !isString(v) {
				return false
			}
			str := v.(string)
			return (str == "basic" || str == "premium")
		}

		valueMap := map[string]interface{}{
			"uid":                  isStringNotEmpty,
			"loan_reference_id":    isStringNotEmpty,
			"user_level":           isUserLevel,
			"loan_level":           isLoanLevel,
			"store_name":           isStringNotEmpty,
			"latest_balance":       isNumeric,
			"phone_number":         isNumeric,
			"requested_at":         isDate,
			"sales_executive":      isSalesExecutiveValid,
			"summary_1_month":      isMap,
			"summary_3_month":      isMap,
			"account_created_at":   isDate,
			"last_transaction_at":  isDate,
			"account_length_month": isNumeric,
			"latest_latitude":      isLatitude,
			"latest_longitude":     isLongitude,
		}
		for k, v := range valueMap {
			if _, ok := profilingData[k]; ok {
				if v.(func(interface{}) bool)(profilingData[k]) == false {
					return k
				}
			} else {
				templateData := map[string]interface{}{"Field": k}
				msg, _ := lang.CurrentTranslation.Translate("ERR_KEY_NOT_FOUND", templateData)
				return msg
			}

		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_PROFILING_DATA", templateData)
		return msg
	},
}

var whitelistIpValidator = &ValidatorWrapper{
	Key: "whitelist_ip_validator",
	ValidateField: func(value interface{}) interface{} {
		isIPAllowed := func(IP string) bool {
			forbiddenIPs := []string{"0.0.0.0"}
			for _, v := range forbiddenIPs {
				if strings.TrimSpace(IP) == v {
					return false
				}
			}
			return true
		}

		isValidIP4 := func(ipAddress string) bool {
			ipAddress = strings.TrimSpace(ipAddress)
			re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
			if re.MatchString(ipAddress) {
				return true
			}
			return false
		}

		whitelistIps := value.([]string)
		for _, ip := range whitelistIps {
			if !isIPAllowed(ip) || !isValidIP4(ip) {
				return ip
			}
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_WHITELIST", templateData)
		return msg
	},
}

var multipleOfValidator = &ValidatorWrapper{
	Key: "multiple_of_validator",
	ValidateField: func(value interface{}) interface{} {
		const MULTIPLE_OF_VALUE float64 = 50000.00
		v := value.(float64)
		result := math.Mod(v, MULTIPLE_OF_VALUE)
		if math.IsNaN(result) {
			return value
		}
		if result != 0 {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(float64)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_MULTIPLY_LOAN_AMOUNT", templateData)
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
		msg, _ := lang.CurrentTranslation.Translate("ERR_ALPHASPACE", templateData)
		return msg
	},
}

var nameValidator = &ValidatorWrapper{
	Key: "name_validator",
	ValidateField: func(value interface{}) interface{} {
		if str, ok := value.(string); !ok {
			return value
		} else {
			str = strings.TrimSpace(str)
			if len(str) == 0 {
				return str
			}
			if re, err := regexp.Compile(`^[a-zA-Z',. ]+$`); err != nil {
				return str
			} else if !re.MatchString(str) {
				return str
			}
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_VALUE", templateData)
		return msg
	},
}

var dateYYYYMMDDValidator = &ValidatorWrapper{
	Key: "date_yyyymmdd_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		_, err := timeutil.Parse(str, timeutil.ISO8601TimeDate)
		if err != nil {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_DATE", templateData)
		return msg
	},
}

var callbackEventTypeValidator = &ValidatorWrapper{
	Key: "callback_event_type_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !enum.IsCallbackEventTypeValid(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_CALLBACK_EVENT_TYPE", templateData)
		return msg
	},
}

var registrationStatusValidator = &ValidatorWrapper{
	Key: "registration_status_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !enum.IsRegistrationStatusValid(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_REGISTRATION_STATUS", templateData)
		return msg
	},
}
var emergencyContactRelationshipValidator = &ValidatorWrapper{
	Key: "emergency_contact_relationship_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !enum.IsEmergencyContactRelationshipValid(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_EMERGENCY_CONTACT", templateData)
		return msg
	},
}
var registrationImageTypeValidator = &ValidatorWrapper{
	Key: "registration_image_type_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !enum.IsRegistrationImageTypeValid(str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_IMAGE_TYPE", templateData)
		return msg
	},
}
var seTypeValidator = &ValidatorWrapper{
	Key: "se_type_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !enum.IsSeTypeValid(&str) {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		failedValues := vw.ValidateField(e.Value())
		templateData := map[string]interface{}{"Field": failedValues.(string)}
		msg, _ := lang.CurrentTranslation.Translate("ERR_SE_TYPE", templateData)
		return msg
	},
}

var passwordValidator = &ValidatorWrapper{
	Key: "password_validator",
	ValidateField: func(value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return value
		}
		if !regexp.MustCompile(`^((?:.*)(?:.*\d)(?:.*[a-zA-Z])|(?:.*[a-zA-Z])(?:.*\d)).{0,}$`).MatchString(str) {
			return value
		}
		if len(str) < 8 || len(str) > 64 {
			return value
		}
		return nil
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		msg, _ := lang.CurrentTranslation.Translate("ERR_PASSWORD")
		return msg
	},
}

var baseLatLongValidator = &ValidatorWrapper{
	Key: "base_lat_long_validator",
	ValidateField: func(v interface{}) interface{} {
		valueMap := v.(map[string]interface{})
		value := valueMap["value"]
		pattern := valueMap["pattern"].(string)

		isLatLongString := func(v interface{}, pattern string) bool {
			if null.IsNil(v) {
				return true
			}
			if str, ok := v.(string); ok {
				str = strings.TrimSpace(str)
				re, _ := regexp.Compile(pattern)
				if len(str) == 0 || re.MatchString(str) {
					return true
				}
				return false
			}
			return false
		}
		isLatLongFloat := func(v interface{}, pattern string) bool {
			if null.IsNil(v) {
				return true
			}
			if fl, ok := v.(float64); ok {
				str := strconv.FormatFloat(fl, 'f', 6, 64)
				return isLatLongString(str, pattern)
			}
			return false
		}
		if isLatLongString(value, pattern) || isLatLongFloat(value, pattern) {
			return nil
		}
		return value
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		msg, _ := lang.CurrentTranslation.Translate("ERR_LATITUDE")
		return msg
	},
}

var latitudeValidator = &ValidatorWrapper{
	Key: "latitude_validator",
	ValidateField: func(v interface{}) interface{} {
		valueMap := map[string]interface{}{
			"value":   v,
			"pattern": `^(\+|-)?(?:90(?:(?:\.0{1,20})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,25})?))$`,
		}
		return baseLatLongValidator.ValidateField(valueMap)
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		msg, _ := lang.CurrentTranslation.Translate("ERR_LATITUDE")
		return msg
	},
}

var longitudeValidator = &ValidatorWrapper{
	Key: "longitude_validator",
	ValidateField: func(v interface{}) interface{} {
		valueMap := map[string]interface{}{
			"value":   v,
			"pattern": `^(\+|-)?(?:180(?:(?:\.0{1,20})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,25})?))$`,
		}
		return baseLatLongValidator.ValidateField(valueMap)
	},
	GetMessage: func(e validator.FieldError, vw ValidatorWrapper) string {
		msg, _ := lang.CurrentTranslation.Translate("ERR_LONGITUDE")
		return msg
	},
}

func (o *ValidatorWrapper) Validate(fl validator.FieldLevel, vw ValidatorWrapper) bool {
	return vw.ValidateField(fl.Field().Interface()) == nil
}
