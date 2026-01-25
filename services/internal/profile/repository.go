package profile

import (
	"context"
)

type Repository interface {
	GetProfile(ctx context.Context, id string) (*Profile, error)
	DeleteProfile(ctx context.Context, id string) error
	StoreProfile(ctx context.Context, id string, profile Profile) (string, error)
	GetOffers(ctx context.Context, profileId string) ([]OfferShort, error)
	GetOffer(ctx context.Context, profileId, id string) (*Offer, error)
	DeleteOffer(ctx context.Context, profileId, id string) error
	StoreOffer(ctx context.Context, profileId string, offer Offer) (string, error)
}
