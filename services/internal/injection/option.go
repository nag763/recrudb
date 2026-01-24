package injection

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

type restMiddleware = func(http.Handler) http.Handler

func asValueProvider(f any, groupName string) fx.Option {
	tag := fmt.Sprintf(`group:"%s"`, groupName)
	return fx.Provide(
		fx.Annotate(
			f,
			fx.ResultTags(tag),
		),
	)
}

func AsMcpMiddlewareProvider(f any) fx.Option {
	return asValueProvider(f, "mcp_middlewares")
}

func AsAgentRouter(f any) fx.Option {
	return asValueProvider(f, "http_agent")
}

func AsRouterProvider(f any) fx.Option {
	return asValueProvider(f, "subrouters")
}

func AsApiMiddlewareProvider(f restMiddleware) fx.Option {
	return asValueProvider(func() restMiddleware { return f }, "api_middlewares")
}
