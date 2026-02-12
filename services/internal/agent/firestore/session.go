package firestore

import (
	"time"

	"google.golang.org/adk/session"
)

type FirestoreSession struct {
	Id              string       `json:"id" firestore:"sessionId"`
	Name            string       `json:"name" firestore:"name"`
	IAppName        string       `json:"appName" firestore:"appName"`
	IUserId         string       `json:"userId" firestore:"userId"`
	IState          StateWrapper `json:"state" firestore:"state"`
	IEvent          EventWrapper `json:"eventWrapper" firestore:"eventWrapper"`
	ILastUpdateTime time.Time    `json:"lastUpdateTime" firestore:"lastUpdateTime,serverTimestamp"`
	CreateTime      time.Time    `json:"createTime" firestore:"createTime,serverTimestamp"`
}

type FirestoreSessionShort struct {
	Id              string    `json:"id" firestore:"sessionId"`
	Name            string    `json:"name" firestore:"name"`
	ILastUpdateTime time.Time `json:"lastUpdateTime" firestore:"lastUpdateTime,serverTimestamp"`
}

func NewFirestoreSession(id, appName, userId string, state map[string]any) FirestoreSession {
	return FirestoreSession{
		Id:       id,
		IAppName: appName,
		IUserId:  userId,
		IState: StateWrapper{
			State: state,
		},
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
	return &fs.IState
}

func (fs FirestoreSession) Events() session.Events {
	return &fs.IEvent
}

func (fs FirestoreSession) LastUpdateTime() time.Time {
	return fs.ILastUpdateTime
}
