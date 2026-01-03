package llm

import (
	"context"
	"fmt"
)

const (
	defaultMaxTokens = 1000
)

// Service provides AI completion capabilities that can be injected into other services
type Service struct {
	provider Provider
}

// NewService creates a new LLM service with the given provider
func NewService(provider Provider) *Service {
	return &Service{
		provider: provider,
	}
}

// Complete sends a completion request with the given messages
func (s *Service) Complete(ctx context.Context, systemPrompt string, userPrompt string) (string, error) {
	if userPrompt == "" {
		return "", fmt.Errorf("user prompt cannot be empty")
	}

	messages := []Message{
		{Role: RoleUser, Content: userPrompt},
	}

	if systemPrompt != "" {
		messages = append([]Message{{Role: RoleSystem, Content: systemPrompt}}, messages...)
	}

	return s.CompleteWithMessages(ctx, messages)
}

// CompleteWithMessages sends a completion request with custom messages
func (s *Service) CompleteWithMessages(ctx context.Context, messages []Message) (string, error) {
	if len(messages) == 0 {
		return "", fmt.Errorf("messages cannot be empty")
	}

	request := CompletionRequest{
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   defaultMaxTokens,
	}

	response, err := s.provider.Complete(ctx, request)
	if err != nil {
		return "", fmt.Errorf("completion failed: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

	return response.Choices[0].Message.Content, nil
}
