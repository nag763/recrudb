package agent

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
)

func getMcpTools(allowedTools []string) (tool.Toolset, error) {

	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.StreamableClientTransport{
			Endpoint: "http://localhost:8001/mcp",
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
