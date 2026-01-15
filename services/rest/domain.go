package function

import "errors"

type GCFRequest struct {
	Content string `json:"content"`
}

type GCFResponse struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Content *string `json:"content,omitempty"`
}

func (r *GCFRequest) validate() error {
	if r.Content == "" {
		return errors.New("content musn't be empty")
	}
	return nil
}

func NewResponseSuccess(content string) GCFResponse {
	return GCFResponse{
		Success: true,
		Content: &content,
	}
}

func NewResponseError(err error) GCFResponse {
	errStr := err.Error()
	return GCFResponse{
		Success: false,
		Error:   &errStr,
	}
}
