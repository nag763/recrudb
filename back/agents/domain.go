package function

type AgentGCFRequest struct {
	Content string `json:"content"`
}

type AgentGCFResponse struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Content *string `json:"content,omitempty"`
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
