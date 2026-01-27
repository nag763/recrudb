package identity

import (
	"context"
	"errors"
	"fmt"
	"nag763/recrudb/api/internal/pkg/utils"
	"net/http"
	"strings"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var getFirebaseAuthClient = sync.OnceValues(func() (*auth.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
})

func extractBearer(req http.Header) (string, error) {
	var authorizationHeader string
	if authorizationHeader = req.Get("Authorization"); authorizationHeader == "" {
		if authorizationHeader = req.Get("X-Serverless-Authorization"); authorizationHeader == "" {
			return "", errors.New("Authorization not found")
		}
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	token = strings.TrimSpace(token)
	if token == "" {
		return "", errors.New("authorization header malformated")
	}
	return token, nil
}

func verifyGoogleIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	err := sync.OnceValue(func() error {
		project := utils.GetEnvOrDefault("PROJECT_NAME", "recrudb")
		return utils.SetEnvIfNotExists("GOOGLE_CLOUD_PROJECT", project)

	})()
	if err != nil {
		return nil, err
	}
	client, err := getFirebaseAuthClient()
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return token, nil
}

func VerifyGoogleIDTokenFromRequest(ctx context.Context, req http.Header) (*auth.Token, error) {
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
	payload, err := verifyGoogleIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	return payload.Subject, nil
}

func GetSubFromRequest(ctx context.Context, req http.Header) (string, error) {
	token, err := extractBearer(req)
	if err != nil {
		return "", err
	}
	return getSub(ctx, token)
}
