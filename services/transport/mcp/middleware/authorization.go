package middleware

import (
	"context"
	"nag763/recrudb/api/internal/identity"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AuthorizationHook(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		if _, err := identity.VerifyGoogleIDTokenFromRequest(ctx, req.Header); err != nil {
			return nil, err
		}
		return next(ctx, req)
	}
}
