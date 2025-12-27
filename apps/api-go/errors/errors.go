package errors

import "github.com/vektah/gqlparser/v2/gqlerror"

// Error codes for standardized client-side error handling
const (
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
	ErrCodeInternal     = "INTERNAL_ERROR"
	ErrCodeBadRequest   = "BAD_REQUEST"
)

// New creates a GraphQL error with a code and message
func New(code, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

// WithField adds a field name to the error (useful for validation errors)
func WithField(err *gqlerror.Error, field string) *gqlerror.Error {
	if err.Extensions == nil {
		err.Extensions = make(map[string]interface{})
	}
	err.Extensions["field"] = field
	return err
}

// WithDetails adds additional details to the error
func WithDetails(err *gqlerror.Error, details map[string]interface{}) *gqlerror.Error {
	if err.Extensions == nil {
		err.Extensions = make(map[string]interface{})
	}
	for key, value := range details {
		err.Extensions[key] = value
	}
	return err
}

// Common error constructors
func NotFound(resource string) *gqlerror.Error {
	return New(ErrCodeNotFound, resource+" not found")
}

func Unauthorized(message string) *gqlerror.Error {
	return New(ErrCodeUnauthorized, message)
}

func ValidationError(field, message string) *gqlerror.Error {
	return WithField(
		New(ErrCodeValidation, message),
		field,
	)
}

func Internal(message string) *gqlerror.Error {
	return New(ErrCodeInternal, message)
}
