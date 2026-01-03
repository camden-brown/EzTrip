package llm

import (
	"fmt"
	"os"
)

const (
	envProvider = "LLM_PROVIDER"

	ProviderXAI    = "xai"
	ProviderGemini = "gemini"
)

// providerFactory maps provider names to their constructor functions
// Providers register themselves via RegisterProvider
var providerFactory = map[string]func() (Provider, error){}

// RegisterProvider registers a provider constructor
func RegisterProvider(name string, factory func() (Provider, error)) {
	providerFactory[name] = factory
}

// NewDefaultService creates an LLM service using environment configuration
// Set LLM_PROVIDER env var to select provider (xai, gemini)
// Defaults to xai if not specified
func NewDefaultService() (*Service, error) {
	providerName := os.Getenv(envProvider)
	if providerName == "" {
		providerName = ProviderXAI
	}

	factory, exists := providerFactory[providerName]
	if !exists {
		return nil, fmt.Errorf("unknown LLM provider: %s (available: %v)", providerName, availableProviders())
	}

	provider, err := factory()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize %s provider: %w", providerName, err)
	}

	return NewService(provider), nil
}

func availableProviders() []string {
	providers := make([]string, 0, len(providerFactory))
	for name := range providerFactory {
		providers = append(providers, name)
	}
	return providers
}
