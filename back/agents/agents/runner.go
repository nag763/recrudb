package agents

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
	"nag763.com/recrudbagents/utils"
)

func createModel(ctx context.Context) (model.LLM, error) {
	return gemini.NewModel(ctx, utils.GetEnvOrDefault("MODEL", "gemini-2.5-flash-lite"), &genai.ClientConfig{
		Backend: genai.BackendVertexAI,
		Project: utils.GetEnvOrDefault("PROJECT_NAME", "recrudb"),
	})
}

type AgentInitializer func(context.Context, model.LLM) (agent.Agent, error)

type AgentRunner struct {
	ctx            context.Context
	runner         runner.Runner
	appName        string
	sessionService session.Service
}

func CreateAgentRunner(agentInitializer AgentInitializer) (*AgentRunner, error) {
	sessionService := session.InMemoryService()
	return CreateAgentRunnerWithSessionService(agentInitializer, sessionService)
}

func CreateAgentRunnerWithSessionService(agentInitializer AgentInitializer, sessionService session.Service) (*AgentRunner, error) {
	ctx := context.Background()
	model, err := createModel(ctx)
	if err != nil {
		return nil, err
	}

	agent, err := agentInitializer(ctx, model)
	if err != nil {
		return nil, err
	}

	runner, err := runner.New(runner.Config{
		Agent:          agent,
		SessionService: sessionService,
		AppName:        agent.Name(),
	})
	if err != nil {
		return nil, err
	}

	return &AgentRunner{
		ctx,
		*runner,
		agent.Name(),
		sessionService,
	}, nil
}

func (r AgentRunner) RunSessionless(input string) (string, error) {

	randomUserId := uuid.NewString()

	userMsg := &genai.Content{
		Parts: []*genai.Part{{Text: input}},
		Role:  string(genai.RoleUser),
	}

	var sb strings.Builder

	request, err := r.sessionService.Create(r.ctx, &session.CreateRequest{
		AppName: r.appName,
		UserID:  randomUserId,
	})
	if err != nil {
		log.Printf("error creating session %v", err)
		return "", err
	}

	sessionId := request.Session.ID()

	defer func() {
		if err := r.sessionService.Delete(r.ctx, &session.DeleteRequest{
			AppName:   r.appName,
			UserID:    randomUserId,
			SessionID: sessionId,
		}); err != nil {
			log.Printf("error while removing stateless call err=%v", err)
		}
	}()

	for event, err := range r.runner.Run(r.ctx, randomUserId, sessionId, userMsg, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		log.Printf("received err=%v response=%v\n", err, event)
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
