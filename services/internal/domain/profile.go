package domain

import "time"

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
