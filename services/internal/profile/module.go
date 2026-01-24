package profile

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("profile",
	fx.Provide(NewProfileRepository),
	fx.Provide(NewProfileTool),
	fx.Provide(NewProfileHandler),
	injection.AsRouterProvider(NewProfileRouter),
)
