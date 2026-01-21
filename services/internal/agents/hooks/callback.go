package hooks

import (
	"nag763/recrudb/api/internal/domain"

	"google.golang.org/adk/tool"
)

func InjectMetadata(ctx tool.Context, tool tool.Tool, args map[string]any) (map[string]any, error) {
	args["metadata"] = domain.Metadata{
		UserId: ctx.UserID(),
	}
	return nil, nil
}
