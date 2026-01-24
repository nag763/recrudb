package profile

import (
	"nag763/recrudb/api/internal/subrouter"

	"go.uber.org/fx"
)

var Module = fx.Module("profile",
	fx.Provide(NewProfileRepository),
	fx.Provide(NewProfileTool),
	fx.Provide(NewProfileHandler),
	subrouter.AsRouterProvider(NewProfileRouter),
)
