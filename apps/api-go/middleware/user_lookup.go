package middleware

import (
	"eztrip/api-go/logger"
	"eztrip/api-go/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserLookupMiddleware creates middleware that extracts Auth0 ID from JWT and stores it in context.
// This must be placed after Auth0JWTMiddleware in the middleware chain.
func UserLookupMiddleware(userService *user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		validated := c.Request.Context().Value(validatedClaimsContextKey{})
		if validated == nil {
			c.Next()
			return
		}

		auth0ID, err := extractSubjectFromClaims(validated)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("Failed to extract subject from validated claims")
			c.Next()
			return
		}

		c.Request = c.Request.WithContext(user.SetUserAuth0ID(c.Request.Context(), auth0ID))
		c.Next()
	}
}
