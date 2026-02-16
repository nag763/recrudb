package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/secret"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/session"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

//go:embed prompts/joboffer_agent.txt
var jobofferAgentPrompt string

func NewJobOfferAgent(agentConfig *AgentConfig, secretService secret.SecretService) (*injection.HttpAgent, error) {
	allowedTools := []string{"stores_offer"}

	mcpTools, err := getMcpTools(allowedTools, secretService.NewAuthenticatedClient())

	jobOfferMatcher, err := llmagent.New(llmagent.Config{
		Name:                "job_offer_agent",
		Description:         "Analyzes a job offer, provide a score and details.",
		Model:               agentConfig.Model,
		Instruction:         jobofferAgentPrompt,
		Toolsets:            []tool.Toolset{mcpTools},
		BeforeToolCallbacks: []llmagent.BeforeToolCallback{injectMetadata},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				FunctionCallingConfig: restrictConfigTo(allowedTools),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &injection.HttpAgent{
		Path:    "/jobOffer",
		Agent:   jobOfferMatcher,
		Session: session.InMemoryService(),
	}, nil
}
