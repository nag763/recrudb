package domain

import "time"

type Communication struct {
	First    time.Time `json:"first`
	Reminder time.Time `json:"reminder"`
	Last     time.Time `json:"last"`
}

func NewCommunication() Communication {
	initial := time.Now()
	reminder := initial.AddDate(0, 0, 7)
	last := initial.AddDate(0, 0, 21)
	return Communication{
		initial,
		reminder,
		last,
	}
}
