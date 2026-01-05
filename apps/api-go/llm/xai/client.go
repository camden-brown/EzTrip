package xai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"eztrip/api-go/llm"
	"eztrip/api-go/logger"

	"github.com/sirupsen/logrus"
)

func init() {
	llm.RegisterProvider(llm.ProviderXAI, func() (llm.Provider, error) {
		return NewClient()
	})
}

const (
	apiURL         = "https://api.x.ai/v1/chat/completions"
	defaultModel   = "grok-4-1-fast-reasoning"
	defaultTimeout = 30 * time.Second
	envAPIKey      = "XAI_API_KEY"
)

// Client implements the Provider interface for xAI/Grok
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new xAI client using the XAI_API_KEY environment variable
func NewClient() (*Client, error) {
	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("XAI_API_KEY environment variable is required")
	}

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}, nil
}

// Complete sends a completion request to xAI and returns the response
func (c *Client) Complete(ctx context.Context, request llm.CompletionRequest) (*llm.CompletionResponse, error) {
	c.setDefaults(&request)

	httpReq, err := c.buildRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var response llm.CompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (c *Client) setDefaults(request *llm.CompletionRequest) {
	if request.Model == "" {
		request.Model = defaultModel
	}
	if request.MaxTokens == 0 {
		request.MaxTokens = 1000
	}
}

func (c *Client) buildRequest(ctx context.Context, request llm.CompletionRequest) (*http.Request, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to send request to xAI API")
		return nil, fmt.Errorf("failed to send request to xAI API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Log.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
			"body":        string(body),
		}).Error("xAI API returned error status")
		return nil, fmt.Errorf("xAI API error (status %d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}
