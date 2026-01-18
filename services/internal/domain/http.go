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

type HttpResponse[T any] struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Content    T      `json:"content,omitempty"`
	StatusCode int    `json:"-"`
}

type HttpStringResponse = HttpResponse[string]

func (r *HttpRequest) validate() error {
	if r.Content == "" {
		return errors.New("content musn't be empty")
	}
	return nil
}

func SuccessResponse[T any](content T) render.Renderer {
	return &HttpResponse[T]{
		Success:    true,
		Content:    content,
		StatusCode: 200,
	}
}

func NoContentResponse() render.Renderer {
	return &HttpStringResponse{
		Success:    true,
		StatusCode: 204,
	}
}

func ErrInternalResponse(err error) render.Renderer {
	return &HttpStringResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 500,
	}
}

func ErrInvalidResponse(err error) render.Renderer {
	return &HttpStringResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 400,
	}
}

func ErrUnauthorizedResponse(err error) render.Renderer {
	return &HttpStringResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: 403,
	}
}

func ErrNotFoundResponse() render.Renderer {
	return &HttpStringResponse{
		Success:    false,
		Error:      "this route doesn't exist",
		StatusCode: 404,
	}

}

func (h *HttpRequest) Bind(r *http.Request) error {
	if err := h.validate(); err != nil {
		return fmt.Errorf("malformed request : %v", err)
	}
	return nil
}

func (h *HttpResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, h.StatusCode)
	return nil
}
