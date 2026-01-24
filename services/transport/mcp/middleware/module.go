package middleware

import (
	"nag763/recrudb/api/internal/middleware"

	"go.uber.org/fx"
)

var Module = fx.Module("middelware_mcp",
	middleware.AsMcpMiddlewareProvider(NewAuthorizationMiddleware),
	middleware.AsMcpMiddlewareProvider(NewLoggingMiddleware),
)
