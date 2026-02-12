package middleware

import (
	"context"
	"nag763/recrudb/api/internal/metadata"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

func NewLoggingMiddleware(log *zap.SugaredLogger) server.ToolHandlerMiddleware {
	return func(thf server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			start := time.Now()
			uuid := uuid.NewString()

			var metadata metadata.MetadataOnly
			if err := req.BindArguments(&metadata); err != nil {
				log.Warn("couldn't retrieve agent on mcp tool call")
			}

			log.Infow("tool call started", "trace", uuid, "params", req.Params.Name, "agent", metadata.Metadata.Agent)

			result, err := thf(ctx, req)

			duration := time.Since(start)

			log.Infow("tool call ended", "trace", uuid, "success", result.IsError || err != nil, "tool", req.Params.Name, "duration", duration, "error", err)
			if result.IsError {
				log.Warnw("tool error detail", "trace", uuid, "content", result.Content)
			}
			return result, err
		}
	}
}
