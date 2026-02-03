package profile

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type FirestoreProfileRepository struct {
	firestoreClient *firestore.Client
}

func NewFirestoreProfileRepository(firestoreClient *firestore.Client) *FirestoreProfileRepository {
	return &FirestoreProfileRepository{
		firestoreClient,
	}
}

func (r *FirestoreProfileRepository) getProfileDocument(id string) *firestore.DocumentRef {
	return r.firestoreClient.Collection("profile").Doc(id)
}

func (r *FirestoreProfileRepository) getProfileDocumentOffers(id string) *firestore.CollectionRef {
	return r.getProfileDocument(id).Collection("offers")
}

func (r *FirestoreProfileRepository) GetProfile(ctx context.Context, id string) (*Profile, error) {
	doc, err := r.getProfileDocument(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var p Profile
	if err := doc.DataTo(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *FirestoreProfileRepository) StoreProfile(ctx context.Context, id string, profile Profile) (string, error) {
	if _, err := r.getProfileDocument(id).Set(ctx, profile); err != nil {
		return "", nil
	}
	return id, nil
}

func (r *FirestoreProfileRepository) DeleteProfile(ctx context.Context, id string) error {
	_, err := r.getProfileDocument(id).Delete(ctx)
	return err
}

func (r *FirestoreProfileRepository) GetOffers(ctx context.Context, id string) ([]OfferShort, error) {
	iter := r.getProfileDocumentOffers(id).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	defer iter.Stop()

	offerShort := []OfferShort{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over documents %w", err)
		}

		var o OfferShort
		if err := doc.DataTo(&o); err != nil {
			return nil, err
		}

		o.Id = doc.Ref.ID
		offerShort = append(offerShort, o)
	}

	return offerShort, nil
}

func (r *FirestoreProfileRepository) GetOffer(ctx context.Context, profileId, id string) (*Offer, error) {
	doc, err := r.getProfileDocumentOffers(profileId).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var o Offer
	if err := doc.DataTo(&o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *FirestoreProfileRepository) DeleteOffer(ctx context.Context, profileId, id string) error {
	_, err := r.getProfileDocumentOffers(profileId).Doc(id).Delete(ctx)
	return err
}

func (r *FirestoreProfileRepository) StoreOffer(ctx context.Context, profileId string, offer Offer) (string, error) {
	uuid := uuid.NewString()
	if _, err := r.getProfileDocumentOffers(profileId).Doc(uuid).Set(ctx, offer); err != nil {
		return "", err
	}
	return uuid, nil
}
