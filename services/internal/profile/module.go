package profile

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("profile",
	fx.Provide(NewProfileRepository),
	fx.Provide(NewProfileHandler),
	injection.AsRouterProvider(NewProfileRouter),
	injection.AsMcpToolProvider(NewGetProfileTool),
	injection.AsMcpToolProvider(NewStoreProfileTool),
)
