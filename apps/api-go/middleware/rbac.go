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

// RequirePermission creates a middleware that checks if the user has permission to access a resource.
// This can be used for REST endpoints, but for GraphQL we'll check in resolvers.
func RequirePermission(resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userUUID := user.GetUserUUID(c.Request.Context())
		if userUUID == "" {
			logger.Log.WithField("resource", resource).Warn("Unauthorized access attempt - no user UUID")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if err := rbac.CheckPermission(c.Request.Context(), userUUID, resource, action); err != nil {
			if err == rbac.ErrUnauthorized {
				logger.Log.WithFields(map[string]interface{}{
					"user_uuid": userUUID,
					"resource":  resource,
					"action":    action,
				}).Warn("Access denied - insufficient permissions")
				c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
			} else {
				logger.Log.WithError(err).Error("Permission check failed")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRole creates a middleware that checks if the user has a specific role.
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userUUID := user.GetUserUUID(c.Request.Context())
		if userUUID == "" {
			logger.Log.WithField("required_role", role).Warn("Unauthorized access attempt - no user UUID")
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

		hasRole, err := rbac.HasRole(enforcer, userUUID, role)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to check role")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		if !hasRole {
			logger.Log.WithFields(map[string]interface{}{
				"user_uuid":     userUUID,
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
// This function retrieves the User UUID from context (set by Auth0 middleware).
func CheckPermissionForGraphQL(ctx context.Context, resource, action string) error {
	userUUID := user.GetUserUUID(ctx)
	if userUUID == "" {
		return rbac.ErrUnauthorized
	}

	return rbac.CheckPermission(ctx, userUUID, resource, action)
}
