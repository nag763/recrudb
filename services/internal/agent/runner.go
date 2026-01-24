package agent

import (
	"context"
	"log"
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

func NewAgentRunner(sessionService session.Service, agent agent.Agent) (*Runner, error) {
	ctx := context.Background()

	runner, err := runner.New(runner.Config{
		Agent:          agent,
		SessionService: sessionService,
		AppName:        agent.Name(),
	})
	if err != nil {
		return nil, err
	}

	return &Runner{
		ctx:            ctx,
		runner:         *runner,
		appName:        agent.Name(),
		sessionService: sessionService,
	}, nil
}

type RunParams struct {
	Input        string
	UserId       string
	ClearSession bool
}

func (r Runner) Run(params RunParams) (string, error) {

	userMsg := &genai.Content{
		Parts: []*genai.Part{{Text: params.Input}},
		Role:  string(genai.RoleUser),
	}

	var sb strings.Builder

	request, err := r.sessionService.Create(r.ctx, &session.CreateRequest{
		AppName: r.appName,
		UserID:  params.UserId,
	})
	if err != nil {
		log.Printf("error creating session %v", err)
		return "", err
	}

	sessionId := request.Session.ID()
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

	for event, err := range r.runner.Run(r.ctx, params.UserId, sessionId, userMsg, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		if err != nil {
			return "", err
		} else if event.Partial {
			for _, p := range event.LLMResponse.Content.Parts {
				sb.WriteString(p.Text)
			}
		}
	}
	return sb.String(), nil
}
