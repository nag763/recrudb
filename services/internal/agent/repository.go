package agent

import "google.golang.org/adk/session"

type AgentRepository struct {
	SessionService session.Service
}

func NewAgentRepository() *AgentRepository {
	return &AgentRepository{
		SessionService: session.InMemoryService(),
	}
}
