package mcp

import (
	"nag763/recrudb/api/transport/mcp/middleware"

	"go.uber.org/fx"
)

var Module = fx.Module("mcp",
	middleware.Module,
	fx.Provide(NewMcpServer),
)
