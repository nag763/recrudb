package clients

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
	"google.golang.org/api/idtoken"

	"nag763/recrudb/api/internal/pkg/utils"
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

type idTokenSource struct {
	TokenSource oauth2.TokenSource
}

func (s *idTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.TokenSource.Token()
	if err != nil {
		return nil, err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("token did not contain an id_token")
	}

	return &oauth2.Token{
		AccessToken: idToken,
		TokenType:   "Bearer",
		Expiry:      token.Expiry,
	}, nil
}

func getTokenSource(ctx context.Context, audience string) (oauth2.TokenSource, error) {
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		if err.Error() != `idtoken: unsupported credentials type` {
			return nil, err
		}
		gts, err := google.DefaultTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		ts = oauth2.ReuseTokenSource(nil, &idTokenSource{TokenSource: gts})
	}
	return ts, nil
}

func getHttpClient(ctx context.Context, endpoint string) (*http.Client, error) {
	tokenSource, err := getTokenSource(ctx, stripPath(endpoint))
	if err != nil {
		return nil, err
	}
	return oauth2.NewClient(ctx, tokenSource), nil
}

func GetMcpTools(ctx context.Context, allowedTools []string) (tool.Toolset, error) {

	endpoint := utils.GetEnvOrDefault(
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
