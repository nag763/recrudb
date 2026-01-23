package agent

import "go.uber.org/fx"

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	fx.Provide(NewAgentRepository),
	fx.Provide(NewAgentService),
	fx.Provide(NewAgentRouter),
)
