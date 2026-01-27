package api

import (
	"nag763/recrudb/api/internal/agent"
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/pkg/logger"
	"nag763/recrudb/api/internal/profile"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	logger.Module,
	db.Module,
	profile.Module,
	agent.Module,

	injection.AsApiMiddlewareProvider(middleware.Logger),
	injection.AsApiMiddlewareProvider(middleware.Recoverer),
	injection.AsApiMiddlewareProvider(corsMiddleware),

	fx.Provide(NewHttpServer),
)
