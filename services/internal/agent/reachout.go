package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/injection"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/session"
)

//go:embed prompts/reachout_agent.txt
var reachOutPrompt string

func NewReachOutAgent(agentConfig *AgentConfig) (*injection.HttpAgent, error) {

	agent, err := llmagent.New(llmagent.Config{
		Name:        "reachout_agent",
		Model:       agentConfig.Model,
		Description: "Help writing cover letters.",
		Instruction: reachOutPrompt,
	})
	if err != nil {
		return nil, err
	}
	return &injection.HttpAgent{
		Agent:   agent,
		Path:    "/reachout",
		Session: session.InMemoryService(),
	}, nil
}
