package domain

type Metadata struct {
	UserId string `json:"userId,omitempty"`
}

type withMetadata[T any] struct {
	Document T        `json:"document"`
	Metadata Metadata `json:"metadata"`
}

type ProfileWithMetadata = withMetadata[Profile]
