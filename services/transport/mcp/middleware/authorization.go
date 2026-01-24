package middleware

import (
	"context"
	"nag763/recrudb/api/internal/pkg/identity"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

func NewAuthorizationMiddleware(log *zap.SugaredLogger) server.ToolHandlerMiddleware {
	return func(thf server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			if _, err := identity.VerifyGoogleIDTokenFromRequest(ctx, req.Header); err != nil {
				log.Warnw("tool authorization rejected", "tool", req.Params.Name)
				return nil, err
			}
			log.Debug("tool authorization accepted")
			return thf(ctx, req)
		}
	}
}
