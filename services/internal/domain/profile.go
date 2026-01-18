package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Profile struct {
	Firstname   string
	Lastname    string
	Contact     Contact `json:",omitempty"`
	Skills      []string
	LookingFor  string
	Experiences []struct {
		StartDate time.Time
		EndDate   time.Time
		Company   string
		Misison   string
		Role      string
		Skills    []string
	}
	Languages []string
	Socials   []struct {
		Name string
		Link string
	}
	Location Address `json:",omitempty"`
}

func GetProfile(ctx context.Context, client *firestore.Client, id string) (*Profile, error) {
	doc, err := client.Collection("profile").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var p Profile
	if err := doc.DataTo(&p); err != nil {
		return nil, err
	}
	return &p, nil
}
