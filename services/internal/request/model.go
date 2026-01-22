package request

import (
	"errors"
	"fmt"
	"net/http"
)

type HttpRequest struct {
	Content string `json:"content"`
}

func (r *HttpRequest) validate() error {
	if r.Content == "" {
		return errors.New("content musn't be empty")
	}
	return nil
}
func (h *HttpRequest) Bind(r *http.Request) error {
	if err := h.validate(); err != nil {
		return fmt.Errorf("malformed request : %v", err)
	}
	return nil
}
