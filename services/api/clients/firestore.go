package clients

import (
	"context"
	"nag763/recrudb/api/utils"
	"sync"

	"cloud.google.com/go/firestore"
)

var GetFirestoreClient = sync.OnceValues(func() (*firestore.Client, error) {
	return firestore.NewClient(context.Background(), utils.GetEnvOrDefault("FIRESTORE_DB", "recrudb"))
})
