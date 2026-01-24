package agent

import (
	"nag763/recrudb/api/internal/subrouter"

	"github.com/go-chi/chi/v5"
)

func NewAgentRouter(agentService *AgentService) subrouter.SubRouter {
	r := chi.NewMux()

	for path, runner := range agentService.Runners {
		r.Post(path, httpHandler(runner))
	}

	return subrouter.SubRouter{
		Path:   "/agents",
		Router: r,
	}
}
