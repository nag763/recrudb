package domain

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type HttpRequest struct {
	Content string `json:"content"`
}

type HttpResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Content    string `json:"content,omitempty"`
	StatusCode int    `json:"-"`
}

func (r *HttpRequest) validate() error {
	if r.Content == "" {
		return errors.New("content musn't be empty")
	}
	return nil
}

func SuccessResponse(content string) render.Renderer {
	return &HttpResponse{
		Success:    true,
		Content:    content,
		StatusCode: 200,
	}
}

func ErrInternalResponse(err error) render.Renderer {
	return &HttpResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 500,
	}
}

func ErrInvalidResponse(err error) render.Renderer {
	return &HttpResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 400,
	}
}

func ErrUnauthorizedResponse(err error) render.Renderer {
	return &HttpResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 403,
	}
}

func ErrNotFoundResponse() render.Renderer {
	return &HttpResponse{
		Success:    false,
		Error:      "this route doesn't exist",
		StatusCode: 404,
	}

}

// Bind satisfies the render.Binder interface
func (h *HttpRequest) Bind(r *http.Request) error {
	if err := h.validate(); err != nil {
		return fmt.Errorf("malformed request : %v", err)
	}
	return nil
}

func (h *HttpResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if h.Error != "" {
		render.Status(r, h.StatusCode)
	}
	return nil
}
