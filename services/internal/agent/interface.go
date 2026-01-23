package agent

import (
	"context"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
)

type AgentInitializer func(ctx context.Context, model model.LLM) (agent.Agent, error)
