package injection

import (
	"fmt"

	"go.uber.org/fx"
)

func asvalueProvider(f any, groupName string) fx.Option {
	tag := fmt.Sprintf(`group:"%s"`, groupName)
	return fx.Provide(
		fx.Annotate(
			f,
			fx.ResultTags(tag),
		),
	)
}

func AsMcpMiddlewareProvider(f any) fx.Option {
	return asvalueProvider(f, "mcp_middlewares")
}

func AsRouterProvider(f any) fx.Option {
	return asvalueProvider(f, "subrouters")
}
