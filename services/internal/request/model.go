package request

import (
	"errors"
	"fmt"
	"io"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/render"
	"google.golang.org/genai"
)

type HttpRequest struct {
	Data         string `json:"data"`
	ClearSession bool   `json:"clearSession,omitempty"`
	SessionId    string `json:"sessionId"`
	ContentType  string `json:"-"`
}

func (r *HttpRequest) validate() error {
	if r.Data == "" {
		return errors.New("data musn't be empty")
	}
	return nil
}

func (h *HttpRequest) Bind(r *http.Request) error {
	if err := h.validate(); err != nil {
		return fmt.Errorf("malformed request : %v", err)
	}
	return nil
}

func FromMultipart(w http.ResponseWriter, r *http.Request) (*HttpRequest, error) {
	file, header, err := r.FormFile("file")
	if err != nil {
		render.Render(w, r, response.InternalResponse(err))
		return nil, err
	}
	defer file.Close()

	mimeType := header.Header.Get("Content-Type")

	fileContent, err := io.ReadAll(file)
	if err != nil {
		render.Render(w, r, response.InternalResponse(err))
		return nil, err
	}

	return &HttpRequest{
		ContentType: mimeType,
		Data:        string(fileContent),
	}, nil
}

func (h HttpRequest) ToGenAiContent() *genai.Content {
	if h.Data == "" {
		return nil
	}
	switch h.ContentType {
	case "":
		return genai.NewContentFromText(h.Data, genai.RoleUser)
	default:
		return genai.NewContentFromBytes([]byte(h.Data), h.ContentType, genai.RoleUser)
	}
}
