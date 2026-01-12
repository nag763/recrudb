package middleware

import (
	"context"
	"errors"
	"fmt"
	"main/utils"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"google.golang.org/api/idtoken"
)

func verifyGoogleIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, utils.GetEnvOrDefault("JWT_AUDIENCE", ""))
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	fmt.Printf("Token verified for user: %v\n", payload.Claims["email"])
	return payload, nil
}

func AuthorizationHook(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader == "" {
			return nil, errors.New("Authorization not found")
		}
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if token == "" {
			return nil, errors.New("Authorization header malformated")
		}
		if _, err := verifyGoogleIDToken(ctx, token); err != nil {
			return nil, errors.New("Token couldn't be verified")
		}
		return next(ctx, req)
	}
}
