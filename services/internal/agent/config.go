package agent

import (
	"context"
	"nag763/recrudb/api/internal/pkg/utils"

	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

type AgentConfig struct {
	Model   model.LLM
	AppName string
}

func restrictConfigTo(tools []string) *genai.FunctionCallingConfig {
	return &genai.FunctionCallingConfig{
		AllowedFunctionNames: tools,
		Mode:                 genai.FunctionCallingConfigModeValidated,
	}
}

func NewAgentConfig() (*AgentConfig, error) {
	model, error := gemini.NewModel(context.Background(), utils.GetEnvOrDefault("MODEL", "gemini-2.5-flash"), &genai.ClientConfig{
		Backend: genai.BackendVertexAI,
		Project: utils.GetEnvOrDefault("PROJECT_NAME", "recrudb"),
	})
	return &AgentConfig{
		Model:   model,
		AppName: utils.GetEnvOrDefault("PROJECT_NAME", "recrudb"),
	}, error
}
