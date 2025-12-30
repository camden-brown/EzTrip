package app

import (
	"os"
	"strings"

	"eztrip/api-go/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureCORS() cors.Config {
	corsConfig := cors.DefaultConfig()

	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",")
	} else {
		if gin.Mode() == gin.DebugMode {
			corsConfig.AllowAllOrigins = true
			logger.Log.WithField("component", "cors").Warn("Allowing all origins in debug mode. Set CORS_ALLOWED_ORIGINS in production.")
		} else {
			logger.Log.WithField("component", "cors").Fatal("CORS_ALLOWED_ORIGINS environment variable is required in production")
		}
	}

	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true

	return corsConfig
}
