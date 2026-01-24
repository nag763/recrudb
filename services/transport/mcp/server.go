package mcp

import (
	"context"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/pkg/utils"
	"nag763/recrudb/api/internal/profile"

	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

type McpParams struct {
	fx.In
	Tools       []injection.Tool               `group:"mcp_tools"`
	Middlewares []server.ToolHandlerMiddleware `group:"mcp_middlewares"`
}

func NewMcpServer(lc fx.Lifecycle, mcpParams McpParams, profileTool *profile.ProfileTool) *server.SSEServer {

	opts := []server.ServerOption{server.WithToolCapabilities(true)}
	for _, middleware := range mcpParams.Middlewares {
		opts = append(opts, server.WithToolHandlerMiddleware(middleware))
	}

	s := server.NewMCPServer(
		"RecruDB Agent Tools",
		"1.0.0",
		opts...,
	)

	for _, tool := range mcpParams.Tools {
		s.AddTool(tool.Tool, tool.Handler)
	}

	httpServer := server.NewSSEServer(s)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go httpServer.Start(utils.GetEnvOrDefault("ADRESS", ":3000"))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})

	return httpServer
}
