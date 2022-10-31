package errors

import (
	goErrors "errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockTranslate(value string) (string, error) {
	return value, nil
}

func mockTranslateFailed(value string) (string, error) {
	return value, goErrors.New("error")
}

func TestE(t *testing.T) {
	type EInput struct {
		HttpCode     int
		MessageCode  int
		Err          error
		WithContextT bool
	}

	type EOutput struct {
		ErrorMessage string
	}

	NewE := func(httpCode, messageCode int, err error, withContextT bool) *EInput {
		return &EInput{httpCode, messageCode, err, withContextT}
	}

	NewEOutput := func(errorMessage string) *EOutput {
		return &EOutput{errorMessage}
	}

	defaultFormatMessage := `{"code":%v,"message":"%v","details":%v}`

	testCases := []struct {
		in        *EInput
		translate func(string) (string, error)
		want      *EOutput
	}{
		{
			in:        NewE(http.StatusUnprocessableEntity, 111, NewGenericError(UNPROCESSABLE_REQUEST), true),
			translate: mockTranslate,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
		{
			in:        NewE(http.StatusUnprocessableEntity, 111, NewGenericError(UNPROCESSABLE_REQUEST), true),
			translate: mockTranslateFailed,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, fmt.Sprintf("error code: %v", NewGenericError(UNPROCESSABLE_REQUEST).(*GenericError).GetCode()), "[]")),
		},
		{
			in:        NewE(http.StatusUnprocessableEntity, 111, NewGenericError(UNPROCESSABLE_REQUEST), false),
			translate: mockTranslate,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, fmt.Sprintf("error code: %v", NewGenericError(UNPROCESSABLE_REQUEST).(*GenericError).GetCode()), "[]")),
		},
		{
			in: NewE(http.StatusUnprocessableEntity, 111, NewGenericErrorWithFn(UNPROCESSABLE_REQUEST, func(_ string) string {
				return "error"
			}), true),
			translate: mockTranslate,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, "error", "[]")),
		},
		{
			in: NewE(http.StatusUnprocessableEntity, 111, NewGenericErrorWithFn(UNPROCESSABLE_REQUEST, func(_ string) string {
				return "error"
			}), true),
			translate: mockTranslateFailed,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, "error", "[]")),
		},
		{
			in:        NewE(http.StatusInternalServerError, 111, goErrors.New("wrong error"), true),
			translate: mockTranslate,
			want:      NewEOutput(fmt.Sprintf(defaultFormatMessage, ERROR_UNKNOWN, "Internal Server Error", `[{"key":"error","value":"wrong error"}]`)),
		},
	}

	for index, tc := range testCases {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		if tc.in.WithContextT {
			c.Set("T", tc.translate)
		}

		c.Status(tc.in.HttpCode)

		E(c, tc.in.HttpCode, tc.in.MessageCode, tc.in.Err)

		assert.Equal(t, 1, len(c.Errors), fmt.Sprintf(`E: errors at index %v`, index))
		assert.Equal(t, tc.in.HttpCode, w.Code, fmt.Sprintf(`E: HttpCode at index %v`, index))
		assert.Equal(t, tc.want.ErrorMessage, w.Body.String(), fmt.Sprintf(`E: response body at index %v`, index))
		assert.True(t, c.IsAborted(), fmt.Sprintf(`E: isAbort() at index %v`, index))
	}
}

func TestErrorCode(t *testing.T) {
	type ErrorCodeInput struct {
		HttpCode     int
		MessageCode  int
		WithContextT bool
	}

	type ErrorCodeOutput struct {
		ErrorMessage string
	}

	NewErrorCode := func(httpCode, messageCode int, withContextT bool) *ErrorCodeInput {
		return &ErrorCodeInput{httpCode, messageCode, withContextT}
	}

	NewErrorCodeOutput := func(errorMessage string) *ErrorCodeOutput {
		return &ErrorCodeOutput{errorMessage}
	}

	defaultFormatMessage := `{"code":%v,"message":"%v","details":%v}`

	testCases := []struct {
		in        *ErrorCodeInput
		translate func(string) (string, error)
		want      *ErrorCodeOutput
	}{
		{
			in:        NewErrorCode(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, true),
			translate: mockTranslate,
			want:      NewErrorCodeOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
		{
			in:        NewErrorCode(http.StatusUnprocessableEntity, 111, true),
			translate: mockTranslate,
			want:      NewErrorCodeOutput(fmt.Sprintf(defaultFormatMessage, 0, fmt.Sprintf("error code: %v", 111), "[]")),
		},
		{
			in:        NewErrorCode(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, true),
			translate: mockTranslateFailed,
			want:      NewErrorCodeOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, fmt.Sprintf("error code: %v", NewGenericError(UNPROCESSABLE_REQUEST).(*GenericError).GetCode()), "[]")),
		},
		{
			in:        NewErrorCode(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, false),
			translate: mockTranslate,
			want:      NewErrorCodeOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, fmt.Sprintf("error code: %v", NewGenericError(UNPROCESSABLE_REQUEST).(*GenericError).GetCode()), "[]")),
		},
	}

	for index, tc := range testCases {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		errorLength := 0
		if tc.in.WithContextT {
			errorLength = 1
			c.Set("T", tc.translate)
		}

		c.Status(tc.in.HttpCode)

		ErrorCode(c, tc.in.HttpCode, tc.in.MessageCode)

		assert.Equal(t, errorLength, len(c.Errors), fmt.Sprintf(`ErrorCode: errors at index %v`, index))
		assert.Equal(t, tc.in.HttpCode, w.Code, fmt.Sprintf(`ErrorCode: HttpCode at index %v`, index))
		assert.Equal(t, tc.want.ErrorMessage, w.Body.String(), fmt.Sprintf(`ErrorCode: response body at index %v`, index))
		assert.True(t, c.IsAborted(), fmt.Sprintf(`ErrorCode: isAbort() at index %v`, index))
	}

}

func TestErrorString(t *testing.T) {
	type ErrorStringInput struct {
		HttpCode     int
		Message      string
		WithContextT bool
	}

	type ErrorStringOutput struct {
		ErrorMessage string
	}

	NewErrorString := func(httpCode int, message string, withContextT bool) *ErrorStringInput {
		return &ErrorStringInput{httpCode, message, withContextT}
	}

	NewErrorStringOutput := func(errorMessage string) *ErrorStringOutput {
		return &ErrorStringOutput{errorMessage}
	}

	defaultFormatMessage := `{"code":%v,"message":"%v","details":%v}`

	testCases := []struct {
		in        *ErrorStringInput
		translate func(string) (string, error)
		want      *ErrorStringOutput
	}{
		{
			in:        NewErrorString(http.StatusUnprocessableEntity, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, true),
			translate: mockTranslate,
			want:      NewErrorStringOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
		{
			in:        NewErrorString(http.StatusUnprocessableEntity, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, true),
			translate: mockTranslate,
			want:      NewErrorStringOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
		{
			in:        NewErrorString(http.StatusUnprocessableEntity, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, true),
			translate: mockTranslateFailed,
			want:      NewErrorStringOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
		{
			in:        NewErrorString(http.StatusUnprocessableEntity, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, false),
			translate: mockTranslate,
			want:      NewErrorStringOutput(fmt.Sprintf(defaultFormatMessage, ERROR_KEYS[UNPROCESSABLE_REQUEST].ErrorCode, ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey, "[]")),
		},
	}

	for index, tc := range testCases {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		if tc.in.WithContextT {
			c.Set("T", tc.translate)
		}

		c.Status(tc.in.HttpCode)

		ErrorString(c, tc.in.HttpCode, tc.in.Message)

		assert.Equal(t, 1, len(c.Errors), fmt.Sprintf(`ErrorString: errors at index %v`, index))
		assert.Equal(t, tc.in.HttpCode, w.Code, fmt.Sprintf(`ErrorString: HttpCode at index %v`, index))
		assert.Equal(t, tc.want.ErrorMessage, w.Body.String(), fmt.Sprintf(`ErrorString: response body at index %v`, index))
		assert.True(t, c.IsAborted(), fmt.Sprintf(`ErrorString: isAbort() at index %v`, index))
	}
}

func TestTranslate(t *testing.T) {
	type TranslateInput struct {
		HttpCode     int
		MessageCode  int
		WithContextT bool
	}

	type TranslateOutput struct {
		ErrorMessage string
	}

	NewTranslate := func(httpCode, messageCode int, withContextT bool) *TranslateInput {
		return &TranslateInput{httpCode, messageCode, withContextT}
	}

	NewTranslateOutput := func(errorMessage string) *TranslateOutput {
		return &TranslateOutput{errorMessage}
	}

	testCases := []struct {
		in        *TranslateInput
		translate func(string) (string, error)
		want      *TranslateOutput
	}{
		{
			in:        NewTranslate(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, true),
			translate: mockTranslate,
			want:      NewTranslateOutput(fmt.Sprintf("%v", ERROR_KEYS[UNPROCESSABLE_REQUEST].MessageKey)),
		},
		{
			in:        NewTranslate(http.StatusUnprocessableEntity, 111, true),
			translate: mockTranslate,
			want:      NewTranslateOutput(fmt.Sprintf("%v", 111)),
		},
		{
			in:        NewTranslate(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, true),
			translate: mockTranslateFailed,
			want:      NewTranslateOutput(fmt.Sprintf("%v", UNPROCESSABLE_REQUEST)),
		},
		{
			in:        NewTranslate(http.StatusUnprocessableEntity, UNPROCESSABLE_REQUEST, false),
			translate: mockTranslate,
			want:      NewTranslateOutput(fmt.Sprintf("%v", UNPROCESSABLE_REQUEST)),
		},
	}

	for index, tc := range testCases {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		if tc.in.WithContextT {
			c.Set("T", tc.translate)
		}
		result := Translate(c, tc.in.MessageCode)

		assert.Equal(t, tc.want.ErrorMessage, result, fmt.Sprintf(`Translate: response body at index %v`, index))
	}
}
