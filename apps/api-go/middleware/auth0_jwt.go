package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"eztrip/api-go/logger"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ErrMissingToken     = errors.New("missing authorization token")
	ErrInvalidToken     = errors.New("invalid token")
	ErrMissingSubClaim  = errors.New("token missing sub claim")
	ErrInvalidIssuerURL = errors.New("invalid issuer URL")
)

type userIDContextKey struct{}

// GetUserIDFromContext retrieves the authenticated user's ID (sub claim) from the request context.
// Returns empty string if not authenticated.
func GetUserIDFromContext(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDContextKey{}).(string); ok {
		return userID
	}
	return ""
}

// Auth0CustomClaims represents custom claims in Auth0 JWT tokens.
type Auth0CustomClaims struct {
	Email string `json:"email"`
}

func (c *Auth0CustomClaims) Validate(_ context.Context) error {
	return nil
}

// Auth0Config holds configuration for Auth0 JWT validation.
type Auth0Config struct {
	IssuerURL string
	Audience  []string
}

// Validate checks if the Auth0 configuration is valid.
func (c Auth0Config) Validate() error {
	if c.IssuerURL == "" {
		return errors.New("AUTH0_ISSUER_URL is required")
	}
	if len(c.Audience) == 0 {
		return errors.New("AUTH0_AUDIENCE is required")
	}
	return nil
}

// normalizeIssuerURL ensures the issuer URL ends with a trailing slash.
func normalizeIssuerURL(issuer string) string {
	if !strings.HasSuffix(issuer, "/") {
		return issuer + "/"
	}
	return issuer
}

// LoadAuth0ConfigFromEnv loads Auth0 configuration from environment variables.
func LoadAuth0ConfigFromEnv() (Auth0Config, error) {
	issuer := strings.TrimSpace(os.Getenv("AUTH0_ISSUER_URL"))
	if issuer == "" {
		return Auth0Config{}, errors.New("AUTH0_ISSUER_URL is required")
	}

	audienceRaw := strings.TrimSpace(os.Getenv("AUTH0_AUDIENCE"))
	if audienceRaw == "" {
		return Auth0Config{}, errors.New("AUTH0_AUDIENCE is required")
	}

	cfg := Auth0Config{
		IssuerURL: normalizeIssuerURL(issuer),
		Audience:  splitAndTrim(audienceRaw, ","),
	}

	if err := cfg.Validate(); err != nil {
		return Auth0Config{}, err
	}

	return cfg, nil
}

// TokenValidator defines the interface for validating JWT tokens.
type TokenValidator interface {
	ValidateToken(ctx context.Context, tokenString string) (interface{}, error)
}

// extractSubjectFromClaims extracts and validates the subject claim from validated token.
func extractSubjectFromClaims(validated interface{}) (string, error) {
	claims, ok := validated.(*validator.ValidatedClaims)
	if !ok || claims == nil {
		return "", ErrInvalidToken
	}

	userID := strings.TrimSpace(claims.RegisteredClaims.Subject)
	if userID == "" {
		return "", ErrMissingSubClaim
	}

	return userID, nil
}

// createTokenValidator creates a new JWT validator with Auth0 configuration.
func createTokenValidator(cfg Auth0Config) (TokenValidator, error) {
	issuerURL, err := url.Parse(cfg.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidIssuerURL, err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		cfg.IssuerURL,
		cfg.Audience,
		validator.WithCustomClaims(func() validator.CustomClaims { return &Auth0CustomClaims{} }),
		validator.WithAllowedClockSkew(60*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return jwtValidator, nil
}

// Auth0JWTMiddleware creates a Gin middleware for Auth0 JWT validation.
func Auth0JWTMiddleware(tokenValidator TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip validation for OPTIONS requests
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// Extract token
		tokenString := extractBearerToken(c.GetHeader("Authorization"))
		if tokenString == "" {
			respondWithError(c, http.StatusUnauthorized, ErrMissingToken)
			return
		}

		// Validate token
		validated, err := tokenValidator.ValidateToken(c.Request.Context(), tokenString)
		if err != nil {
			logTokenValidationError(c, err)
			respondWithError(c, http.StatusUnauthorized, ErrInvalidToken)
			return
		}

		// Extract subject claim
		userID, err := extractSubjectFromClaims(validated)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, err)
			return
		}

		// Store user ID in context
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), userIDContextKey{}, userID))
		c.Next()
	}
}

// Auth0JWTFromEnv creates Auth0 JWT middleware from environment variables.
func Auth0JWTFromEnv() (gin.HandlerFunc, error) {
	cfg, err := LoadAuth0ConfigFromEnv()
	if err != nil {
		return nil, err
	}

	tokenValidator, err := createTokenValidator(cfg)
	if err != nil {
		return nil, err
	}

	return Auth0JWTMiddleware(tokenValidator), nil
}

// respondWithError sends a standardized error response.
func respondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error":   "unauthorized",
		"message": err.Error(),
	})
	c.Abort()
}

// logTokenValidationError logs token validation failures.
func logTokenValidationError(c *gin.Context, err error) {
	logger.Log.WithFields(logrus.Fields{
		"error":  err.Error(),
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Warn("JWT validation failed")
}

func extractBearerToken(authHeader string) string {
	authHeader = strings.TrimSpace(authHeader)
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		return ""
	}
	if !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

func splitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}
