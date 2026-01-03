package utils

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
)

func GetMcpTools(allowedTools []string) (tool.Toolset, error) {
	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.StreamableClientTransport{
			Endpoint: GetEnvOrDefault("MCP_ENDPOINT", ":3000"),
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
