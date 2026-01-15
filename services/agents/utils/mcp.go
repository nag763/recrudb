package utils

import (
	"context"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/oauth2"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
	"google.golang.org/api/idtoken"
)

func stripPath(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""
	return u.String()
}

func GetMcpTools(ctx context.Context, allowedTools []string) (tool.Toolset, error) {

	endpoint := GetEnvOrDefault(
		"MCP_ENDPOINT",
		"http://localhost:3000/sse",
	)
	tokenSource, err := idtoken.NewTokenSource(ctx, stripPath(endpoint))
	if err != nil {
		return nil, err
	}

	httpClient := oauth2.NewClient(ctx, tokenSource)

	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.SSEClientTransport{
			Endpoint:   endpoint,
			HTTPClient: httpClient,
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
