package middleware

import (
	"context"
	"main/utils"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AuthorizationHook(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		bearer, err := utils.ExtractBearer(req)
		if err != nil {
			return nil, err
		}

		if _, err := utils.VerifyGoogleIDToken(ctx, bearer); err != nil {
			return nil, err
		}
		return next(ctx, req)
	}
}
