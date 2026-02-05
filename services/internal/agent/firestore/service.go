package firestore

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/adk/session"
	"google.golang.org/api/iterator"
)

type FirestoreSessionService struct {
	client *firestore.Client
}

type fss = FirestoreSessionService

func NewFirestoreSessionService(client *firestore.Client) *FirestoreSessionService {
	return &FirestoreSessionService{
		client,
	}
}

func (f *fss) getProfileSessions(userId, appName string) (*firestore.CollectionRef, error) {
	if userId == "" {
		return nil, errors.New("user id is nil")
	}
	if appName == "" {
		return nil, errors.New("app name is nil")
	}

	doc := f.client.Collection("profile").Doc(userId)
	if doc == nil {
		return nil, errors.New("the user doesn't exist")
	}
	return doc.Collection("apps").Doc(appName).Collection("sessions"), nil
}

func (f *fss) getSession(userId, appName, sessionId string) (*firestore.DocumentRef, error) {
	if sessionId == "" {
		return nil, errors.New("session id is nil")
	}
	sessions, err := f.getProfileSessions(userId, appName)
	if err != nil {
		return nil, err
	}
	return sessions.Doc(sessionId), nil
}

func (f *fss) Create(ctx context.Context, r *session.CreateRequest) (*session.CreateResponse, error) {
	var sessionId = r.SessionID
	if sessionId == "" {
		sessionId = uuid.NewString()
	}

	doc, err := f.getSession(r.UserID, r.AppName, sessionId)
	if err != nil {
		return nil, err
	}

	fs := NewFirestoreSession(sessionId, r.AppName, r.UserID)

	if _, err = doc.Set(ctx, fs); err != nil {
		return nil, err
	}

	return &session.CreateResponse{
		Session: fs,
	}, nil
}

func (f *fss) Get(ctx context.Context, r *session.GetRequest) (*session.GetResponse, error) {
	doc, err := f.getSession(r.UserID, r.AppName, r.SessionID)
	if err != nil {
		return nil, err
	}

	docRef, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	var fs FirestoreSession
	if err := docRef.DataTo(&fs); err != nil {
		return nil, err
	}

	return &session.GetResponse{
		Session: fs,
	}, nil
}

func (f *fss) List(ctx context.Context, r *session.ListRequest) (*session.ListResponse, error) {
	docs, err := f.getProfileSessions(r.UserID, r.AppName)
	if err != nil {
		return nil, err
	}
	iter := docs.Documents(ctx)
	defer iter.Stop()
	sessions := []session.Session{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var s FirestoreSession
		if err := doc.DataTo(&s); err != nil {
			return nil, err
		}

		sessions = append(sessions, s)
	}

	return &session.ListResponse{
		Sessions: sessions,
	}, nil
}

func (f *fss) Delete(ctx context.Context, r *session.DeleteRequest) error {
	doc, err := f.getSession(r.UserID, r.AppName, r.SessionID)
	if err != nil {
		return err
	}

	if _, err := doc.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func (f *fss) AppendEvent(ctx context.Context, s session.Session, e *session.Event) error {
	doc, err := f.getSession(s.UserID(), s.AppName(), s.ID())
	if err != nil {
		return err
	}

	_, err = doc.Update(ctx, []firestore.Update{
		{
			Path:  "eventWrapper.events",
			Value: firestore.ArrayUnion(e),
		},
	})
	return err
}
