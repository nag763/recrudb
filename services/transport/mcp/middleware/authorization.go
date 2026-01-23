package middleware

import (
	"context"
	"nag763/recrudb/api/internal/pkg/identity"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

type AuthorizationMiddleware struct {
	log *zap.SugaredLogger
}

func NewAuthorizationMiddleware(log *zap.SugaredLogger) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{log}
}

func (m *AuthorizationMiddleware) NewToolHandlerMiddleware(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		if _, err := identity.VerifyGoogleIDTokenFromRequest(ctx, req.Header); err != nil {
			m.log.Warnw("tool authorization rejected", "tool", req.Params.Name)
			return nil, err
		}
		m.log.Debug("tool authorization accepted")
		return next(ctx, req)
	}
}
