package middleware

import (
	"eztrip/api-go/logger"
	"eztrip/api-go/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserLookupMiddleware creates middleware that fetches the User UUID from database using Auth0 ID.
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

		foundUser, err := userService.GetByAuth0ID(c.Request.Context(), auth0ID)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"auth0_id": auth0ID,
				"error":    err.Error(),
			}).Warn("Failed to fetch user by Auth0 ID")
			c.Next()
			return
		}

		c.Request = c.Request.WithContext(user.SetUserUUID(c.Request.Context(), foundUser.ID))
		c.Next()
	}
}
