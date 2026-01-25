package profile

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type ProfileRepository struct {
	firestoreClient *firestore.Client
}

func NewProfileRepository(firestoreClient *firestore.Client) *ProfileRepository {
	return &ProfileRepository{
		firestoreClient,
	}
}

func (r *ProfileRepository) getProfileDocument(id string) *firestore.DocumentRef {
	return r.firestoreClient.Collection("profile").Doc(id)
}

func (r *ProfileRepository) getProfileDocumentOffers(id string) *firestore.CollectionRef {
	return r.getProfileDocument(id).Collection("offers")
}

func (r *ProfileRepository) GetProfile(ctx context.Context, id string) (*Profile, error) {
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

func (r *ProfileRepository) StoreProfile(ctx context.Context, id string, profile Profile) (*firestore.WriteResult, error) {
	return r.getProfileDocument(id).Set(ctx, profile)
}

func (r *ProfileRepository) GetOffers(ctx context.Context, id string) (*[]OfferShort, error) {
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

		offerName, err := doc.DataAt("Name")
		if err != nil {
			offerName = ""
		}
		name, ok := offerName.(string)
		if !ok {
			name = ""
		}

		offerShort = append(offerShort, OfferShort{
			Id:        doc.Ref.ID,
			CreatedAt: doc.CreateTime,
			Name:      name,
		})
	}

	return &offerShort, nil
}

func (r *ProfileRepository) StoreOffer(ctx context.Context, profileId string, offer Offer) (*firestore.WriteResult, error) {
	return r.getProfileDocumentOffers(profileId).Doc(uuid.NewString()).Set(ctx, offer)
}
