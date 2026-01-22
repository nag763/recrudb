package db

import (
	"context"
	"nag763/recrudb/api/internal/utils"

	"cloud.google.com/go/firestore"
	"go.uber.org/fx"
)

func NewFirestoreClient(lc fx.Lifecycle) (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, utils.GetEnvOrDefault("FIRESTORE_DB", "recrudb"))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
	return client, err
}
