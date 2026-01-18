package callback

import "google.golang.org/adk/tool"

func InferUserId(ctx tool.Context, tool tool.Tool, args map[string]any) (map[string]any, error) {
	args["userId"] = ctx.UserID()
	return args, nil
}
