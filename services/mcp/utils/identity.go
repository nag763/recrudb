package utils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/mark3labs/mcp-go/mcp"
)

func extractBearer(req mcp.CallToolRequest) (string, error) {
	authorizationHeader := req.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("Authorization not found")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if token == "" {
		return "", errors.New("Authorization header malformated")
	}
	return token, nil
}

func verifyGoogleIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, GetEnvOrDefault("JWT_AUDIENCE", ""))
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return payload, nil
}

func VerifyGoogleIDTokenFromRequest(ctx context.Context, request mcp.CallToolRequest) (*idtoken.Payload, error) {
	bearer, err := extractBearer(request)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return verifyGoogleIDToken(ctx, bearer)
}

func GetClaims(ctx context.Context, idToken string) (map[string]interface{}, error) {
	payload, err := verifyGoogleIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return payload.Claims, nil
}

func getSub(ctx context.Context, idToken string) (string, error) {
	claims, err := GetClaims(ctx, idToken)
	if err != nil {
		return "", err
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("sub couldn't be extracted")
	}
	return sub, nil
}

func GetSubFromRequest(ctx context.Context, req mcp.CallToolRequest) (string, error) {
	token, err := extractBearer(req)
	if err != nil {
		return "", err
	}
	return getSub(ctx, token)
}
