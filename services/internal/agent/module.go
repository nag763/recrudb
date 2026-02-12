package agent

import (
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/profile"

	"go.uber.org/fx"
)

var Module = fx.Module("agent",
	fx.Provide(NewAgentConfig),
	fx.Provide(firestore.NewFirestoreSessionService),
	fx.Provide(profile.NewFirestoreProfileRepository),
	injection.AsAgentRouter(NewCoordinatorAgent),
	injection.AsAgentRouter(NewCvAnalyzer),
	injection.AsAgentRouter(NewJobOfferAgent),
	injection.AsAgentRouter(NewReachOutAgent),
	injection.AsRouterProvider(NewAgentRouter),
)
