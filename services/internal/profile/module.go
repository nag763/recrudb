package profile

import "go.uber.org/fx"

var Module = fx.Module("profile",
	fx.Provide(NewProfileRepository),
	fx.Provide(NewProfileTool),
	fx.Provide(NewProfileHandler),
)
