package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/secret"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

//go:embed prompts/reachout_agent.txt
var reachOutPrompt string

func NewReachOutAgent(model model.LLM, secretService secret.SecretService) (agent.Agent, error) {

	allowedTools := []string{"save_profile"}
	tools, err := getMcpTools(allowedTools, secretService.NewAuthenticatedClient())
	if err != nil {
		return nil, err
	}

	agent, err := llmagent.New(llmagent.Config{
		Name:        "reachout_agent",
		Model:       model,
		Description: "Help writing cover letters.",
		Instruction: reachOutPrompt,
		Toolsets:    []tool.Toolset{tools},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				FunctionCallingConfig: restrictConfigTo(allowedTools),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return agent, nil
}
