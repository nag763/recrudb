package api

import (
	"nag763/recrudb/api/internal/agent"
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/pkg/logger"
	"nag763/recrudb/api/internal/profile"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
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
	injection.AsApiMiddlewareProvider(middleware.ThrottleBacklog(10, 50, time.Second*10)),
	injection.AsApiMiddlewareProvider(middleware.Timeout(time.Second*75)),
	injection.AsApiMiddlewareProvider(httprate.LimitByRealIP(100, 1*time.Minute)),

	fx.Provide(NewHttpServer),
)
