package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/viper"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
)

func GetMcpTools(allowedTools []string) (tool.Toolset, error) {
	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.StreamableClientTransport{
			Endpoint: viper.GetString("MCP_ENDPOINT"),
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
