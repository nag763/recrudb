package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Profile struct {
	Firstname   string
	Lastname    string   `json:",omitempty"`
	Contact     Contact  `json:",omitempty"`
	Skills      []string `json:",omitempty"`
	LookingFor  string   `json:",omitempty"`
	Experiences []struct {
		StartDate time.Time `json:",omitempty"`
		EndDate   time.Time `json:",omitempty"`
		Company   string    `json:",omitempty"`
		Misison   string    `json:",omitempty"`
		Role      string    `json:",omitempty"`
		Skills    []string  `json:",omitempty"`
	}
	Languages []string `json:",omitempty"`
	Socials   []struct {
		Name string `json:",omitempty"`
		Link string `json:",omitempty"`
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
