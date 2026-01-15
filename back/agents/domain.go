package function

import "errors"

type AgentGCFRequest struct {
	Content string `json:"content"`
}

type AgentGCFResponse struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Content *string `json:"content,omitempty"`
}

func (r *AgentGCFRequest) validate() error {
	if r.Content == "" {
		return errors.New("content musn't be empty")
	}
	return nil
}

func NewResponseSuccess(content string) AgentGCFResponse {
	return AgentGCFResponse{
		Success: true,
		Content: &content,
	}
}

func NewResponseError(err error) AgentGCFResponse {
	errStr := err.Error()
	return AgentGCFResponse{
		Success: false,
		Error:   &errStr,
	}
}
