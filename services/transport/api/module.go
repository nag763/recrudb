package api

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	fx.Provide(
		NewHttpServer,
	),
	fx.Invoke(func(*chi.Mux) {}),
)
