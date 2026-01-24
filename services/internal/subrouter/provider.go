package subrouter

import (
	"go.uber.org/fx"
)

func asRouter(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(`group:"subrouters"`),
	)
}

func AsRouterProvider(f any) fx.Option {
	return fx.Provide(asRouter(f))
}
