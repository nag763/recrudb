package agent

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	fx.Provide(NewAgentService),
	injection.AsRouterProvider(NewAgentRouter),
)
