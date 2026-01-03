package utils

import (
	"context"
	"strings"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

const (
	userID  = "user1234"
	appName = "Google Search_agent"
)

func CallAgent(ctx context.Context, a agent.Agent, prompt string) (string, error) {
	// TODO : Pass in lazy
	sessionService := session.InMemoryService()
	session, err := sessionService.Create(ctx, &session.CreateRequest{
		AppName: appName,
		UserID:  userID,
	})
	if err != nil {
		return "", err
	}

	config := runner.Config{
		AppName:        appName,
		Agent:          a,
		SessionService: sessionService,
	}
	r, err := runner.New(config)
	if err != nil {
		return "", err
	}

	sessionID := session.Session.ID()
	userMsg := &genai.Content{
		Parts: []*genai.Part{{Text: prompt}},
		Role:  string(genai.RoleUser),
	}

	var sb strings.Builder
	for event, err := range r.Run(ctx, userID, sessionID, userMsg, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		if err != nil {
			return "", err
		} else if event.Partial {
			for _, p := range event.LLMResponse.Content.Parts {
				sb.WriteString(p.Text)
			}
		}
	}
	return sb.String(), nil
}
