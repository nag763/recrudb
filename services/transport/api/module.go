package api

import (
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	injection.AsApiMiddlewareProvider(middleware.Logger),
	injection.AsApiMiddlewareProvider(middleware.Recoverer),
	fx.Provide(NewHttpServer),
)
