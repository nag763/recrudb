package profile

import (
	"nag763/recrudb/api/internal/address"
	"nag763/recrudb/api/internal/contact"
	"nag763/recrudb/api/internal/metadata"
	"time"
)

type Profile struct {
	Firstname   string
	Lastname    string          `json:",omitempty"`
	Contact     contact.Contact `json:",omitempty"`
	Skills      []string        `json:",omitempty"`
	LookingFor  string          `json:",omitempty"`
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
	Location address.Address `json:",omitempty"`
}

type ProfileMcp = metadata.McpDocument[Profile]
type ProfileMcpWithMeta = metadata.WithMetadata[Profile]
