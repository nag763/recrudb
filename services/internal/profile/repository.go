package profile

import (
	"context"

	"cloud.google.com/go/firestore"
)

type ProfileRepository struct {
	firestoreClient *firestore.Client
}

func NewProfileRepository(firestoreClient *firestore.Client) *ProfileRepository {
	return &ProfileRepository{
		firestoreClient,
	}
}

func (r *ProfileRepository) GetProfile(ctx context.Context, id string) (*Profile, error) {
	doc, err := r.firestoreClient.Collection("profile").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var p Profile
	if err := doc.DataTo(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProfileRepository) StoreProfile(ctx context.Context, id string, profile Profile) (*firestore.WriteResult, error) {
	return r.firestoreClient.Collection("profile").Doc(id).Set(ctx, profile)
}
