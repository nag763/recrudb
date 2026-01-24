package agent

import (
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5"
)

func NewAgentRouter(agentService *AgentService) injection.SubRouter {
	r := chi.NewMux()

	for path, runner := range agentService.Runners {
		r.Post(path, httpHandler(runner))
	}

	return injection.SubRouter{
		Path:   "/agents",
		Router: r,
	}
}
