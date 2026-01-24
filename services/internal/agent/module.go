package agent

import (
	"nag763/recrudb/api/internal/subrouter"

	"go.uber.org/fx"
)

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	fx.Provide(NewAgentService),
	subrouter.AsRouterProvider(NewAgentRouter),
)
