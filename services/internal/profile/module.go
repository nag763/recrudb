package profile

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("profile",
	fx.Provide(
		fx.Annotate(NewFirestoreProfileRepository,
			fx.As(new(Repository)),
		),
	),
	fx.Provide(NewProfileHandler),
	injection.AsRouterProvider(NewProfileRouter),
	injection.AsMcpToolProvider(NewGetProfileTool),
	injection.AsMcpToolProvider(NewStoreProfileTool),
	injection.AsMcpToolProvider(NewStoreOfferTool),
	injection.AsMcpToolProvider(NewGetOfferTool),
)
