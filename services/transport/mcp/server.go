package mcp

import (
	"context"
	"nag763/recrudb/api/internal/offer"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/internal/utils"
	"nag763/recrudb/api/transport/mcp/middleware"

	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

func NewMcpServer(lc fx.Lifecycle, profileTool *profile.ProfileTool, offerTool *offer.OfferTool) *server.SSEServer {

	s := server.NewMCPServer(
		"RecruDB Agent Tools",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithToolHandlerMiddleware(middleware.AuthorizationHook),
	)

	s.AddTool(profileTool.NewGetProfileTool())
	s.AddTool(profileTool.NewStoreProfileTool())

	s.AddTool(offerTool.NewStoreOfferTool())

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
