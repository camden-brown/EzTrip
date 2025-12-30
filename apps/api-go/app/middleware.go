package app

import (
	"eztrip/api-go/logger"
	"eztrip/api-go/middleware"
	"eztrip/api-go/rbac"
	"eztrip/api-go/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMiddleware(router *gin.Engine, db *gorm.DB, enforcer *casbin.Enforcer) error {
	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ErrorHandler())
	router.Use(cors.New(ConfigureCORS()))

	auth0JWT, err := middleware.Auth0JWTFromEnv()
	if err != nil {
		return err
	}
	router.Use(auth0JWT)

	userService := user.NewService(db)
	router.Use(middleware.UserLookupMiddleware(userService))
	router.Use(middleware.RBACMiddleware(enforcer))

	return nil
}

func InitializeRBAC(db *gorm.DB) (*casbin.Enforcer, error) {
	enforcer, err := rbac.NewEnforcer(db)
	if err != nil {
		return nil, err
	}
	logger.Log.WithField("component", "rbac").Info("Casbin RBAC enforcer initialized")

	if err := rbac.InitializePolicies(enforcer); err != nil {
		return nil, err
	}
	
	policyCount := len(rbac.DefaultPolicies)
	logger.Log.WithFields(map[string]interface{}{
		"component":     "rbac",
		"policy_count":  policyCount,
	}).Info("RBAC policies initialized")

	return enforcer, nil
}
