package agent

import (
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	fx.Provide(firestore.NewFirestoreSessionService),
	injection.AsAgentRouter(NewCoordinatorAgent),
	injection.AsAgentRouter(NewCvAnalyzer),
	injection.AsAgentRouter(NewJobOfferAgent),
	injection.AsRouterProvider(NewAgentRouter),
)
