package middleware

import (
	"context"
	"net/http"

	"eztrip/api-go/logger"
	"eztrip/api-go/rbac"
	"eztrip/api-go/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// RBACMiddleware creates a Gin middleware that injects the Casbin enforcer into the context.
// The enforcer automatically loads policies and role assignments from the database.
func RBACMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := rbac.SetEnforcerInContext(c.Request.Context(), enforcer)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// RequireRole creates a middleware that checks if the user has a specific role.
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth0ID := user.GetUserAuth0ID(c.Request.Context())
		if auth0ID == "" {
			logger.Log.WithField("required_role", role).Warn("Unauthorized access attempt - no user Auth0 ID")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		enforcer, err := rbac.GetEnforcerFromContext(c.Request.Context())
		if err != nil {
			logger.Log.WithError(err).Error("Failed to get enforcer from context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		hasRole, err := rbac.HasRole(enforcer, auth0ID, role)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to check role")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		if !hasRole {
			logger.Log.WithFields(map[string]interface{}{
				"user_auth0_id": auth0ID,
				"required_role": role,
			}).Warn("Access denied - missing required role")
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CheckPermissionForGraphQL is a helper function to check permissions in GraphQL resolvers.
// It automatically checks both direct permissions and role-based permissions via Casbin.
// This function retrieves the User Auth0 ID from context (set by Auth0 middleware).
func CheckPermissionForGraphQL(ctx context.Context, resource, action string) error {
	auth0ID := user.GetUserAuth0ID(ctx)
	if auth0ID == "" {
		return rbac.ErrUnauthorized
	}

	return rbac.CheckPermission(ctx, auth0ID, resource, action)
}
