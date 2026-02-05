package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/secret"

	"google.golang.org/adk/agent/llmagent"
)

//go:embed prompts/coordinator.txt
var coordinatorPrompt string

func NewCoordinatorAgent(agentConfig *AgentConfig, secretService secret.SecretService, firestoreSession *firestore.FirestoreSessionService) (*injection.HttpAgent, error) {

	agent, err := llmagent.New(llmagent.Config{
		Name:        "coordinator",
		Model:       agentConfig.Model,
		Description: "Coordinates and reply to general knowledge questions.",
		Instruction: coordinatorPrompt,
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
