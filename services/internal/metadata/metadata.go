package metadata

type Metadata struct {
	UserId string `json:"userId,omitempty"`
	Agent  string `json:"agent,omitempty"`
}

type McpDocument[T any] struct {
	Document T
}

type WithMetadata[T any] struct {
	McpDocument[T]
	Metadata Metadata `json:"metadata"`
}

type MetadataOnly = WithMetadata[struct{}]
