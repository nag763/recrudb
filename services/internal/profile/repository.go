package profile

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Repository interface {
	GetProfile(ctx context.Context, id string) (*Profile, error)
	StoreProfile(ctx context.Context, id string, profile Profile) (*firestore.WriteResult, error)
	GetOffers(ctx context.Context, id string) (*[]OfferShort, error)
	StoreOffer(ctx context.Context, profileId string, offer Offer) (*firestore.WriteResult, error)
}
