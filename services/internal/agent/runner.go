package agent

import (
	"context"
	"errors"
	"log"
	"nag763/recrudb/api/internal/injection"
	"strings"

	"go.uber.org/zap"
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
	logger         *zap.SugaredLogger
}

func NewAgentRunner(agent *injection.HttpAgent, logger *zap.SugaredLogger) (*Runner, error) {
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
		logger:         logger,
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
	r.logger.Debugf("runner %s received %+v", r.appName, params)
	var sb strings.Builder

	var sessionId string
	if params.SessionId == "" {
		request, err := r.sessionService.Create(r.ctx, &session.CreateRequest{
			AppName: r.appName,
			UserID:  params.UserId,
		})
		if err != nil {
			r.logger.Errorf("error creating session %v", err)
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

RunnerLoop:
	for event, err := range r.runner.Run(r.ctx, params.UserId, sessionId, params.Input, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		if err != nil {
			var apiError genai.APIError
			switch {
			case err.Error() == "TODO: last event is not final":
				r.logger.Warnf("runner %s: non-final event detected", r.appName)
				break RunnerLoop
			case errors.As(err, &apiError):
				r.logger.Errorf("genai api error: %v", err)
				return nil, err
			default:
				return nil, err
			}
		} else if event.Partial {
			for _, p := range event.LLMResponse.Content.Parts {
				sb.WriteString(p.Text)
			}
		}
	}
	r.logger.Infof("runner %s ran with success", r.appName)
	return &RunResponse{
		Content:   sb.String(),
		SessionId: sessionId,
	}, nil
}
