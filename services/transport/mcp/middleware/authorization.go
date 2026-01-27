package middleware

import (
	"context"
	"nag763/recrudb/api/internal/secret"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

func NewAuthorizationMiddleware(log *zap.SugaredLogger, secretService secret.SecretService) server.ToolHandlerMiddleware {
	return func(thf server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			if ok := secretService.VerifySecretFromHeader(&req.Header); !ok {
				return mcp.NewToolResultError("shared secret is not recognized"), nil
			}
			log.Debug("tool authorization accepted")
			return thf(ctx, req)
		}
	}
}
