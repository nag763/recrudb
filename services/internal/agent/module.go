package agent

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	injection.AsAgentRouter(NewCvAnalyzer),
	injection.AsAgentRouter(NewJobOfferAgent),
	injection.AsRouterProvider(NewAgentRouter),
)
