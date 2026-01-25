package profile

import (
	"nag763/recrudb/api/internal/metadata"
	"time"
)

type Contact struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"number,omitempty"`
}

type Communication struct {
	First    time.Time `json:"first"`
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

type OfferShort struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name,omitempty" firestore:"name"`
}

type Offer struct {
	EnterpriseName        string    `json:"enterpriseName"`
	Contact               Contact   `json:"contact"`
	Address               Address   `json:"address"`
	TechnicalRequirements []string  `json:"techRequirements"`
	CandidateScore        float32   `json:"candidateScore"`
	HowToApply            string    `json:"howToApply"`
	Description           string    `json:"description"`
	AgentRecap            string    `json:"agentRecap"`
	Name                  string    `json:"name"`
	CreatedAt             time.Time `json:"-" firestore:"CreatedAt,serverTimestamp"`
}

type Address struct {
	AddressLine1 string `json:"line1,omitempty"`
	AddressLine2 string `json:"line2,omitempty"`
	AddressLine3 string `json:"line3,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
}

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
	Location  Address   `json:",omitempty"`
	CreatedAt time.Time `json:"-" firestore:"CreatedAt,serverTimestamp"`
}

type ProfileMcp = metadata.McpDocument[Profile]
type ProfileMcpWithMeta = metadata.WithMetadata[Profile]
