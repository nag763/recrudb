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

func storeWithTimestamp(ctx context.Context, docRef *firestore.DocumentRef, data interface{}) (*firestore.WriteResult, error) {
	if _, err := docRef.Set(ctx, data); err != nil {
		return nil, err
	}
	return docRef.Set(ctx, map[string]interface{}{"CreatedAt": firestore.ServerTimestamp}, firestore.MergeAll)
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
	docRef := r.getProfileDocument(id)
	return storeWithTimestamp(ctx, docRef, profile)
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
			return nil, fmt.Errorf("failed to iterate over documents")
		}

		offerName, err := doc.DataAt("Name")
		if err != nil {
			offerName = ""
		}

		offerShort = append(offerShort, OfferShort{
			Id:        doc.Ref.ID,
			CreatedAt: doc.CreateTime,
			Name:      offerName.(string),
		})
	}

	return &offerShort, nil
}

func (r *ProfileRepository) StoreOffer(ctx context.Context, id string, offer Offer) (*firestore.WriteResult, error) {
	docRef := r.getProfileDocumentOffers(id).Doc(uuid.NewString())
	return storeWithTimestamp(ctx, docRef, offer)
}
