package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Client handles Auth0 Management API operations
type Client struct {
	domain       string
	clientID     string
	clientSecret string
	audience     string
	httpClient   *http.Client
	accessToken  string
	tokenExpiry  time.Time
}

const (
	tokenExpiryBuffer    = 300 * time.Second // 5 minute buffer
	defaultTimeout       = 10 * time.Second
	headerContentType    = "Content-Type"
	headerAuthorization  = "Authorization"
	contentTypeJSON      = "application/json"
	defaultConnection    = "Username-Password-Authentication"
	pathOAuthToken       = "/oauth/token"
	pathUsers            = "/api/v2/users"
	envAuth0Domain       = "AUTH0_DOMAIN"
	envAuth0ClientID     = "AUTH0_CLIENT_ID"
	envAuth0ClientSecret = "AUTH0_CLIENT_SECRET"
	envAuth0Connection   = "AUTH0_CONNECTION"
)

// NewClient creates a new Auth0 Management API client
func NewClient() (*Client, error) {
	domain := strings.TrimSpace(os.Getenv(envAuth0Domain))
	clientID := strings.TrimSpace(os.Getenv(envAuth0ClientID))
	clientSecret := strings.TrimSpace(os.Getenv(envAuth0ClientSecret))

	if domain == "" || clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("AUTH0_DOMAIN, AUTH0_CLIENT_ID, and AUTH0_CLIENT_SECRET are required")
	}

	return &Client{
		domain:       domain,
		clientID:     clientID,
		clientSecret: clientSecret,
		audience:     fmt.Sprintf("https://%s/api/v2/", domain),
		httpClient:   &http.Client{Timeout: defaultTimeout},
	}, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// getAccessToken retrieves a Management API access token
func (c *Client) getAccessToken() error {
	// Return cached token if still valid
	if c.accessToken != "" && time.Now().Before(c.tokenExpiry) {
		return nil
	}

	payload := map[string]string{
		"client_id":     c.clientID,
		"client_secret": c.clientSecret,
		"audience":      c.audience,
		"grant_type":    "client_credentials",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal token request: %w", err)
	}

	url := fmt.Sprintf("https://%s/oauth/token", c.domain)
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to request access token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := readResponseBody(resp)
		if err != nil {
			return fmt.Errorf("auth0 token request failed with status %d: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("auth0 token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	c.accessToken = tokenResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn)*time.Second - tokenExpiryBuffer)

	return nil
}

// buildURL constructs a full URL with the given path
func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("https://%s%s", c.domain, path)
}

// readResponseBody reads and returns the response body with error handling
func readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return body, nil
}

// CreateUserRequest represents the request to create a user in Auth0
type CreateUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
	FirstName  string `json:"given_name,omitempty"`
	LastName   string `json:"family_name,omitempty"`
}

// Auth0User represents a user response from Auth0
type Auth0User struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
}

// CreateUser creates a new user in Auth0
func (c *Client) CreateUser(email, password, firstName, lastName string) (*Auth0User, error) {
	if err := c.getAccessToken(); err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	connection := getAuth0Connection()

	req := CreateUserRequest{
		Email:      email,
		Password:   password,
		Connection: connection,
		FirstName:  firstName,
		LastName:   lastName,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create user request: %w", err)
	}

	resp, err := c.makeAuthenticatedRequest(http.MethodPost, c.buildURL(pathUsers), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create user in auth0: %w", err)
	}
	defer resp.Body.Close()

	body, err := readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("auth0 create user failed with status %d: %s", resp.StatusCode, string(body))
	}

	var auth0User Auth0User
	if err := json.Unmarshal(body, &auth0User); err != nil {
		return nil, fmt.Errorf("failed to decode auth0 user response: %w", err)
	}

	return &auth0User, nil
}

// getAuth0Connection returns the Auth0 connection to use, with fallback to default
func getAuth0Connection() string {
	connection := os.Getenv(envAuth0Connection)
	if connection == "" {
		return defaultConnection
	}
	return connection
}

// DeleteUser deletes a user from Auth0 by user ID
func (c *Client) DeleteUser(userID string) error {
	if err := c.getAccessToken(); err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	url := c.buildURL(fmt.Sprintf("%s/%s", pathUsers, userID))
	resp, err := c.makeAuthenticatedRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user from auth0: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, err := readResponseBody(resp)
		if err != nil {
			return fmt.Errorf("auth0 delete user failed with status %d: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("auth0 delete user failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// makeAuthenticatedRequest creates and executes an authenticated HTTP request
func (c *Client) makeAuthenticatedRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set(headerAuthorization, "Bearer "+c.accessToken)
	if body != nil {
		req.Header.Set(headerContentType, contentTypeJSON)
	}

	return c.httpClient.Do(req)
}
