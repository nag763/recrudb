package agent

import (
	"nag763/recrudb/api/internal/metadata"

	"google.golang.org/adk/tool"
)

func injectMetadata(ctx tool.Context, tool tool.Tool, args map[string]any) (map[string]any, error) {
	args["metadata"] = metadata.Metadata{
		UserId: ctx.UserID(),
	}
	return nil, nil
}
