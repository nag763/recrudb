package offer

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

type OfferRepository struct {
	firestoreClient *firestore.Client
}

func NewOfferRepository(firestoreClient *firestore.Client) *OfferRepository {
	return &OfferRepository{
		firestoreClient,
	}
}

func (r *OfferRepository) StoreOffer(ctx context.Context, offer Offer) (*firestore.WriteResult, error) {
	uuid := uuid.NewString()
	return r.firestoreClient.Collection("offers").Doc(uuid).Set(ctx, offer)
}
