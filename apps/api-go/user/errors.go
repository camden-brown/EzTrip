package user

import (
	appErrors "eztrip/api-go/errors"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	ErrCodeDuplicateEmail = "USER_DUPLICATE_EMAIL"
)

// DuplicateEmailError returns an error for when an email is already in use
func DuplicateEmailError() *gqlerror.Error {
	return appErrors.WithField(
		appErrors.New(ErrCodeDuplicateEmail, "Email already in use"),
		"email",
	)
}
