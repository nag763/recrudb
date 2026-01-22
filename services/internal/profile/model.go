package profile

import (
	"nag763/recrudb/api/internal/domain"
	"time"
)

type Profile struct {
	Firstname   string
	Lastname    string         `json:",omitempty"`
	Contact     domain.Contact `json:",omitempty"`
	Skills      []string       `json:",omitempty"`
	LookingFor  string         `json:",omitempty"`
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
	Location domain.Address `json:",omitempty"`
}

type ProfileMcp = domain.McpDocument[Profile]
type ProfileMcpWithMeta = domain.WithMetadata[Profile]
