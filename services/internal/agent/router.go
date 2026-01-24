package agent

import (
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"google.golang.org/adk/session"
)

type AgentRouterParams struct {
	fx.In

	HttpAgents []*injection.HttpAgent `group:"http_agent"`
}

func NewAgentRouter(agentConfig *AgentConfig, agentRouterParams AgentRouterParams) (injection.SubRouter, error) {
	r := chi.NewMux()
	sessionService := session.InMemoryService()

	for _, ha := range agentRouterParams.HttpAgents {
		runner, err := NewAgentRunner(sessionService, ha.Agent)
		if err != nil {
			return injection.SubRouter{}, err
		}
		r.Post(ha.Path, httpHandler(runner))
	}

	return injection.SubRouter{
		Path:   "/agents",
		Router: r,
	}, nil
}
