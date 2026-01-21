package domain

type Metadata struct {
	UserId string `json:"userId,omitempty"`
}

type mcpDocument[T any] struct {
	Document T
}

type withMetadata[T any] struct {
	mcpDocument[T]
	Metadata Metadata `json:"metadata"`
}
