package middleware

import (
	"net/http"

	"travel-app/api-go/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
}

// ErrorHandler catches panics and returns proper error responses
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.WithFields(logrus.Fields{
					"error":  err,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("Panic recovered")

				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "internal_server_error",
					Message: "An unexpected error occurred",
					Path:    c.Request.URL.Path,
				})
				c.Abort()
			}
		}()

		c.Next()

		// Handle errors added to context
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			logger.Log.WithFields(logrus.Fields{
				"error":  err.Err.Error(),
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			}).Error("Request error")

			// Return first error
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "bad_request",
				Message: err.Err.Error(),
				Path:    c.Request.URL.Path,
			})
		}
	}
}
