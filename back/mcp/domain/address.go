package domain

type Address struct {
	AddressLine1 string `json:"line1,omitempty"`
	AddressLine2 string `json:"line2,omitempty"`
	AddressLine3 string `json:"line3,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
}
