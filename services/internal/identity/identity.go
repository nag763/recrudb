package identity

import (
	"context"
	"errors"
	"fmt"
	"nag763/recrudb/api/internal/utils"
	"strings"

	"cloud.google.com/go/auth/credentials/idtoken"
)

type HeaderAccesser interface {
	Get(string) string
}

func extractBearer(req HeaderAccesser) (string, error) {
	authorizationHeader := req.Get("Authorization")
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
	payload, err := idtoken.Validate(ctx, idToken, utils.GetEnvOrDefault("JWT_AUDIENCE", ""))
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return payload, nil
}

func VerifyGoogleIDTokenFromRequest(ctx context.Context, req HeaderAccesser) (*idtoken.Payload, error) {
	bearer, err := extractBearer(req)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return verifyGoogleIDToken(ctx, bearer)
}

func getClaims(ctx context.Context, idToken string) (map[string]interface{}, error) {
	payload, err := verifyGoogleIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return payload.Claims, nil
}

func getSub(ctx context.Context, idToken string) (string, error) {
	claims, err := getClaims(ctx, idToken)
	if err != nil {
		return "", err
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("sub couldn't be extracted")
	}
	return sub, nil
}

func GetSubFromRequest(ctx context.Context, req HeaderAccesser) (string, error) {
	token, err := extractBearer(req)
	if err != nil {
		return "", err
	}
	return getSub(ctx, token)
}
