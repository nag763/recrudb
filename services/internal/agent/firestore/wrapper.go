package firestore

import (
	"iter"

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

type StateWrapper struct {
	State map[string]any `json:"state" firestore:"state"`
}

func (sw StateWrapper) Get(key string) (any, error) {
	val, ok := sw.State[key]
	if !ok {
		return nil, session.ErrStateKeyNotExist
	}
	return val, nil
}

func (sw StateWrapper) Set(key string, value any) error {
	sw.State[key] = value
	return nil
}

func (sw StateWrapper) All() iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for key, val := range sw.State {
			if !yield(key, val) {
				return
			}
		}
	}
}
