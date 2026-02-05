package firestore

import (
	"iter"
	"time"

	"google.golang.org/adk/session"
)

type EventWrapper struct {
	Events []session.Event `firestore:"events"`
}

func (ew EventWrapper) All() iter.Seq[*session.Event] {
	return func(yield func(*session.Event) bool) {
		for i := range ew.Events {
			if !yield(&(ew.Events[i])) {
				return
			}
		}
	}
}

func (ew EventWrapper) Len() int {
	return len(ew.Events)
}

func (ew *EventWrapper) At(i int) *session.Event {
	if i < 0 || i >= ew.Len() {
		return nil
	}
	return &ew.Events[i]
}

type FirestoreSession struct {
	Id              string        `firestore:"sessionId"`
	IAppName        string        `firestore:"appName"`
	IUserId         string        `firestore:"userId"`
	IState          session.State `firestore:"-"`
	IEvent          EventWrapper  `firestore:"eventWrapper"`
	ILastUpdateTime time.Time     `firestore:"time,serverTimestamp"`
}

func NewFirestoreSession(id, appName, userId string) FirestoreSession {
	return FirestoreSession{
		Id:       id,
		IAppName: appName,
		IUserId:  userId,
	}
}

func (fs FirestoreSession) ID() string {
	return fs.Id
}
func (fs FirestoreSession) AppName() string {
	return fs.IAppName
}

func (fs FirestoreSession) UserID() string {
	return fs.IUserId
}

func (fs FirestoreSession) State() session.State {
	return fs.IState
}

func (fs FirestoreSession) Events() session.Events {
	return &fs.IEvent
}

func (fs FirestoreSession) LastUpdateTime() time.Time {
	return fs.ILastUpdateTime
}
