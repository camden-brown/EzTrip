package validation

import (
	"fmt"
	"regexp"
	"strings"

	appErrors "eztrip/api-go/errors"

	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	validate       *validator.Validate
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	numberRegex    = regexp.MustCompile(`[0-9]`)
)

func init() {
	validate = validator.New()

	validate.RegisterValidation("password_complexity", validatePasswordComplexity)
}

func validatePasswordComplexity(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	return uppercaseRegex.MatchString(password) &&
		lowercaseRegex.MatchString(password) &&
		numberRegex.MatchString(password)
}

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return appErrors.ValidationError("", "Invalid input")
	}

	for _, fieldError := range validationErrors {
		return convertValidationError(fieldError)
	}

	return appErrors.ValidationError("", "Invalid input")
}

func convertValidationError(fieldError validator.FieldError) *gqlerror.Error {
	field := strings.ToLower(string(fieldError.Field()[0])) + fieldError.Field()[1:]
	message := getErrorMessage(fieldError)

	return appErrors.WithField(
		appErrors.New(appErrors.ErrCodeValidation, message),
		field,
	)
}

func getErrorMessage(fieldError validator.FieldError) string {
	field := fieldError.Field()

	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
	case "password_complexity":
		return "Password must contain at least one uppercase letter, one lowercase letter, and one number"
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
