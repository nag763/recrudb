package agent

import (
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
)

func getMcpTools(allowedTools []string, client *http.Client) (tool.Toolset, error) {

	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.StreamableClientTransport{
			Endpoint:   "http://localhost:8001/mcp",
			HTTPClient: client,
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
