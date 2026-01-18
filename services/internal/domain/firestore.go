package domain

type firestoreDocumentMCPRequest[T any] struct {
	Document T      `json:"document"`
	UserId   string `json:"userId"`
}

type ProfileDocumentMCPRequest = firestoreDocumentMCPRequest[Profile]
