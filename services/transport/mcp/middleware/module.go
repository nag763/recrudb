package middleware

import (
	"nag763/recrudb/api/internal/injection"

	"go.uber.org/fx"
)

var Module = fx.Module("middelware_mcp",
	injection.AsMcpMiddlewareProvider(NewLoggingMiddleware),
)
