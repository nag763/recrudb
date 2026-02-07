package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/secret"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
)

//go:embed prompts/coordinator.txt
var coordinatorPrompt string

func NewCoordinatorAgent(agentConfig *AgentConfig, secretService secret.SecretService, firestoreSession *firestore.FirestoreSessionService) (*injection.HttpAgent, error) {

	allowedTools := []string{"get_profile"}
	tools, err := getMcpTools(allowedTools, secretService.NewAuthenticatedClient())
	if err != nil {
		return nil, err
	}

	jobOfferAgent, err := NewJobOfferAgent(agentConfig, secretService)
	if err != nil {
		return nil, err
	}
	jobOfferTool := agenttool.New(jobOfferAgent.Agent, &agenttool.Config{})

	agent, err := llmagent.New(llmagent.Config{
		Name:                "coordinator",
		Model:               agentConfig.Model,
		Description:         "Coordinates and reply to general knowledge questions.",
		Instruction:         coordinatorPrompt,
		Tools:               []tool.Tool{jobOfferTool},
		Toolsets:            []tool.Toolset{tools},
		BeforeToolCallbacks: []llmagent.BeforeToolCallback{injectMetadata},
	})
	if err != nil {
		return nil, err
	}

	return &injection.HttpAgent{
		Path:    "/coordinator",
		Agent:   agent,
		Session: firestoreSession,
	}, nil
}
