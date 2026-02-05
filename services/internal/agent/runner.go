package agent

import (
	"context"
	"log"
	"nag763/recrudb/api/internal/injection"
	"strings"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

type Runner struct {
	ctx            context.Context
	runner         runner.Runner
	appName        string
	sessionService session.Service
}

func NewAgentRunner(agent *injection.HttpAgent) (*Runner, error) {
	ctx := context.Background()

	runner, err := runner.New(runner.Config{
		Agent:          agent.Agent,
		SessionService: agent.Session,
		AppName:        agent.Agent.Name(),
	})
	if err != nil {
		return nil, err
	}

	return &Runner{
		ctx:            ctx,
		runner:         *runner,
		appName:        agent.Agent.Name(),
		sessionService: agent.Session,
	}, nil
}

type RunParams struct {
	Input        *genai.Content
	UserId       string
	SessionId    string
	ClearSession bool
}

type RunResponse struct {
	Content   string `json:"content"`
	SessionId string `json:"sessionId"`
}

func (r Runner) Run(params RunParams) (*RunResponse, error) {
	var sb strings.Builder

	var sessionId string
	if params.SessionId == "" {
		request, err := r.sessionService.Create(r.ctx, &session.CreateRequest{
			AppName: r.appName,
			UserID:  params.UserId,
		})
		if err != nil {
			log.Printf("error creating session %v", err)
			return nil, err
		}
		sessionId = request.Session.ID()
	} else {
		req, err := r.sessionService.Get(r.ctx, &session.GetRequest{
			AppName:   r.appName,
			UserID:    params.UserId,
			SessionID: params.SessionId,
		})
		if err != nil {
			return nil, err
		}
		sessionId = req.Session.ID()
	}

	if params.ClearSession {
		defer func() {
			if err := r.sessionService.Delete(r.ctx, &session.DeleteRequest{
				AppName:   r.appName,
				UserID:    params.UserId,
				SessionID: sessionId,
			}); err != nil {
				log.Printf("error while removing stateless call err=%v", err)
			}
		}()
	}

	for event, err := range r.runner.Run(r.ctx, params.UserId, sessionId, params.Input, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		if err != nil {
			return nil, err
		} else if event.Partial {
			for _, p := range event.LLMResponse.Content.Parts {
				sb.WriteString(p.Text)
			}
		}
	}
	return &RunResponse{
		Content:   sb.String(),
		SessionId: sessionId,
	}, nil
}
