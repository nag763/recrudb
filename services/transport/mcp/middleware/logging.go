package middleware

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

type LoggingMiddleware struct {
	log *zap.SugaredLogger
}

func NewLoggingMiddleware(log *zap.SugaredLogger) *LoggingMiddleware {
	return &LoggingMiddleware{log}
}

func (m *LoggingMiddleware) ToolMiddleware(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		start := time.Now()
		uuid := uuid.NewString()
		m.log.Info("tool call started", "trace", uuid, "params", req.Params.Name)

		result, err := next(ctx, req)

		duration := time.Since(start)
		m.log.Infow("tool call ended", "trace", uuid, "success", err != nil, "tool", req.Params.Name, "duration", duration, "error=", err)

		return result, err
	}
}
