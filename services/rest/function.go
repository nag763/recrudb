package function

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetProfile", GetProfile)
}

var getFirestoreClient = sync.OnceValues(func() (*firestore.Client, error) {
	return firestore.NewClient(context.Background(), GetEnvOrDefault("FIRESTORE_DB", "recrudb"))
})

func writeError(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(NewResponseError(err))
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sub, err := GetSubFromRequest(ctx, r)
	if err != nil {
		log.Printf("error while trying to access sub %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	client, err := getFirestoreClient()
	if err != nil {
		writeError(w, err)
		return
	}

	dsnap, err := client.Collection("profile").Doc(sub).Get(ctx)
	if err != nil {
		writeError(w, errors.New("error while retrieving document"))
		return
	}

	if err := json.NewEncoder(w).Encode(dsnap); err != nil {
		writeError(w, errors.New("error while serializing response"))
		return
	}
}
