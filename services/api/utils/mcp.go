package utils

import (
	"context"
	"net/http"
	"net/url"
	"strings"

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

func getHttpClient(ctx context.Context, endpoint string) (*http.Client, error) {
	if strings.Contains(endpoint, "localhost") {
		return nil, nil
	}

	tokenSource, err := idtoken.NewTokenSource(ctx, stripPath(endpoint))
	if err != nil {
		return nil, err
	}
	return oauth2.NewClient(ctx, tokenSource), nil
}

func GetMcpTools(ctx context.Context, allowedTools []string) (tool.Toolset, error) {

	endpoint := GetEnvOrDefault(
		"MCP_ENDPOINT",
		"http://localhost:3000/sse",
	)

	httpClient, err := getHttpClient(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.SSEClientTransport{
			Endpoint:   endpoint,
			HTTPClient: httpClient,
		},
		ToolFilter: tool.StringPredicate(allowedTools),
	})

	return toolset, err
}
