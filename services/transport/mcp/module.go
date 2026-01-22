package mcp

import (
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

var Module = fx.Module("mcp",
	fx.Provide(NewMcpServer),
	fx.Invoke(func(*server.SSEServer) {}),
)
