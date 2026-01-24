package middleware

import (
	"go.uber.org/fx"
)

func asMcpMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(`group:"mcp_middlewares"`),
	)
}

func AsMcpMiddlewareProvider(f any) fx.Option {
	return fx.Provide(asMcpMiddleware(f))
}
