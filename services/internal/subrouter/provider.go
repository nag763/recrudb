package subrouter

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type SubRouter struct {
	Path   string
	Router chi.Router
}

func asRouter(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(`group:"subrouters"`),
	)
}

func AsRouterProvider(f any) fx.Option {
	return fx.Provide(asRouter(f))
}
