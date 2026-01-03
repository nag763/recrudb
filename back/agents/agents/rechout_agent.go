package agents

import (
	_ "embed"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
)

//go:embed prompts/reachout_agent.txt
var reachOutPrompt string

func NewReachOutAgent(model model.LLM) (agent.Agent, error) {

	agent, err := llmagent.New(llmagent.Config{
		Name:        "reachout_agent",
		Model:       model,
		Description: "Writes motivation mails.",
		Instruction: reachOutPrompt,
	})
	if err != nil {
		return nil, err
	}
	return agent, nil
}
