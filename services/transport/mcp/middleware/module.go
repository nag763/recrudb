package middleware

import "go.uber.org/fx"

var Module = fx.Module("middelware_mcp",
	fx.Provide(NewAuthorizationMiddleware),
	fx.Provide(NewLoggingMiddleware),
	fx.Provide(NewMiddlewareRepository),
)
