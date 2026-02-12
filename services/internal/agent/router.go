package agent

import (
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AgentRouterParams struct {
	fx.In

	HttpAgents []*injection.HttpAgent `group:"http_agent"`
}

func NewAgentRouter(agentConfig *AgentConfig, firestoreSessionService *firestore.FirestoreSessionService, agentRouterParams AgentRouterParams, logger *zap.SugaredLogger) (injection.SubRouter, error) {
	r := chi.NewMux()

	for _, ha := range agentRouterParams.HttpAgents {
		runner, err := NewAgentRunner(ha, logger)
		if err != nil {
			return injection.SubRouter{}, err
		}
		r.Post(ha.Path, interactionHttpHandler(runner))
		r.Get(ha.Path+"/sessions", sessionListHttpHandler(runner, firestoreSessionService))
		r.Get(ha.Path+"/sessions/{sessionId}", sessionGetHttpHandler(runner, firestoreSessionService))
		r.Delete(ha.Path+"/sessions/{sessionId}", sessionDeleteHttpHandler(runner, firestoreSessionService))
	}

	return injection.SubRouter{
		Path:   "/agents",
		Router: r,
	}, nil
}
