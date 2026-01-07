package rbac

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	appErrors "eztrip/api-go/errors"
	"eztrip/api-go/logger"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrUnauthorized     = errors.New("unauthorized: insufficient permissions")
	ErrEnforcerNotFound = errors.New("enforcer not found in context")
)

var DefaultPolicies = [][]string{
	{"admin", "*", "*"},
	{"user", "users", "read"},
	{"user", "user", "read"},
	{"user", "currentUser", "*"},
}

type enforcerContextKey struct{}

// GetEnforcerFromContext retrieves the Casbin enforcer from the request context.
func GetEnforcerFromContext(ctx context.Context) (*casbin.Enforcer, error) {
	if enforcer, ok := ctx.Value(enforcerContextKey{}).(*casbin.Enforcer); ok {
		return enforcer, nil
	}
	return nil, ErrEnforcerNotFound
}

// SetEnforcerInContext stores the Casbin enforcer in the request context.
func SetEnforcerInContext(ctx context.Context, enforcer *casbin.Enforcer) context.Context {
	return context.WithValue(ctx, enforcerContextKey{}, enforcer)
}

// NewEnforcer creates a new Casbin enforcer using the official GORM adapter.
// This stores both policies and role assignments in the database.
func NewEnforcer(db *gorm.DB) (*casbin.Enforcer, error) {
	modelPath, err := getModelPath()
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"error":     err.Error(),
		}).Error("Failed to get RBAC model path")
		return nil, fmt.Errorf("failed to get model path: %w", err)
	}

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"error":     err.Error(),
		}).Error("Failed to create GORM adapter for Casbin")
		return nil, fmt.Errorf("failed to create GORM adapter: %w", err)
	}

	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component":  "rbac",
			"model_path": modelPath,
			"error":      err.Error(),
		}).Error("Failed to create Casbin enforcer")
		return nil, fmt.Errorf("failed to create enforcer: %w", err)
	}

	enforcer.EnableAutoSave(true)

	return enforcer, nil
}

// getModelPath returns the absolute path to the model file.
func getModelPath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get current file path")
	}

	dir := filepath.Dir(filename)
	modelPath := filepath.Join(dir, "rbac_model.conf")

	return modelPath, nil
}

// CheckPermission checks if a user has permission to perform an action on a resource.
// This enforces both direct permissions and role-based permissions.
func CheckPermission(ctx context.Context, userID, resource, action string) error {
	enforcer, err := GetEnforcerFromContext(ctx)
	if err != nil {
		return err
	}

	allowed, err := enforcer.Enforce(userID, resource, action)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"resource":  resource,
			"action":    action,
			"error":     err.Error(),
		}).Error("Permission check failed")
		return fmt.Errorf("permission check failed: %w", err)
	}

	if !allowed {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"resource":  resource,
			"action":    action,
		}).Warn("Permission denied")
		return ErrUnauthorized
	}

	return nil
}

// InitializePolicies sets up the initial policy rules in the database.
// This should be called once during application setup if the casbin_rule table is empty.
// Supports wildcards (*) for resources and actions.
func InitializePolicies(enforcer *casbin.Enforcer) error {
	existingPolicies, _ := enforcer.GetPolicy()
	if len(existingPolicies) > 0 {
		logger.Log.WithFields(map[string]interface{}{
			"component":      "rbac",
			"existing_count": len(existingPolicies),
		}).Debug("Policies already exist, skipping initialization")
		return nil
	}

	_, err := enforcer.AddPolicies(DefaultPolicies)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component":    "rbac",
			"policy_count": len(DefaultPolicies),
			"error":        err.Error(),
		}).Error("Failed to add default policies")
		return fmt.Errorf("failed to add policies: %w", err)
	}

	logger.Log.WithFields(map[string]interface{}{
		"component":    "rbac",
		"policy_count": len(DefaultPolicies),
	}).Info("Default policies added to database")

	return nil
}

// AddRoleForUser adds a role assignment for a user (auto-saved to DB).
func AddRoleForUser(enforcer *casbin.Enforcer, userID, role string) error {
	if _, err := enforcer.AddRoleForUser(userID, role); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"role":      role,
			"error":     err.Error(),
		}).Error("Failed to add role for user")
		return fmt.Errorf("failed to add role: %w", err)
	}
	logger.Log.WithFields(map[string]interface{}{
		"component": "rbac",
		"user_id":   userID,
		"role":      role,
	}).Info("Role assigned to user")
	return nil
}

// RemoveRoleForUser removes a role assignment from a user (auto-saved to DB).
func RemoveRoleForUser(enforcer *casbin.Enforcer, userID, role string) error {
	if _, err := enforcer.DeleteRoleForUser(userID, role); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"role":      role,
			"error":     err.Error(),
		}).Error("Failed to remove role from user")
		return fmt.Errorf("failed to remove role: %w", err)
	}
	logger.Log.WithFields(map[string]interface{}{
		"component": "rbac",
		"user_id":   userID,
		"role":      role,
	}).Info("Role removed from user")
	return nil
}

// UpdateUserRole updates a user's role by removing all existing roles and adding the new one.
func UpdateUserRole(enforcer *casbin.Enforcer, userID, newRole string) error {
	if _, err := enforcer.DeleteRolesForUser(userID); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"error":     err.Error(),
		}).Error("Failed to remove existing roles from user")
		return fmt.Errorf("failed to remove existing roles: %w", err)
	}

	if _, err := enforcer.AddRoleForUser(userID, newRole); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"new_role":  newRole,
			"error":     err.Error(),
		}).Error("Failed to add new role to user")
		return fmt.Errorf("failed to add new role: %w", err)
	}

	logger.Log.WithFields(map[string]interface{}{
		"component": "rbac",
		"user_id":   userID,
		"new_role":  newRole,
	}).Info("User role updated")

	return nil
}

// GetRolesForUser returns all roles assigned to a user.
func GetRolesForUser(enforcer *casbin.Enforcer, userID string) ([]string, error) {
	roles, err := enforcer.GetRolesForUser(userID)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"user_id":   userID,
			"error":     err.Error(),
		}).Error("Failed to get roles for user")
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}
	return roles, nil
}

// HasRole checks if a user has a specific role.
func HasRole(enforcer *casbin.Enforcer, userID, role string) (bool, error) {
	return enforcer.HasRoleForUser(userID, role)
}

// GetUsersForRole returns all User UUIDs with a specific role.
func GetUsersForRole(enforcer *casbin.Enforcer, role string) ([]string, error) {
	users, err := enforcer.GetUsersForRole(role)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"component": "rbac",
			"role":      role,
			"error":     err.Error(),
		}).Error("Failed to get users for role")
		return nil, fmt.Errorf("failed to get users for role: %w", err)
	}
	return users, nil
}

// RequireAdminRole checks if the user has the admin role
func RequireAdminRole(ctx context.Context, userUUID uuid.UUID) error {
	enforcer, err := GetEnforcerFromContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to get RBAC enforcer: %w", err)
	}

	hasRole, err := HasRole(enforcer, userUUID.String(), "admin")
	if err != nil {
		return fmt.Errorf("failed to check user role: %w", err)
	}

	if !hasRole {
		logger.Log.WithFields(logrus.Fields{
			"user_id": userUUID.String(),
			"role":    "admin",
		}).Warn("Unauthorized admin access attempt")
		return appErrors.Forbidden("You do not have permission to perform this action")
	}

	return nil
}
